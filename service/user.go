package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"shopping-cart/builder"
	"shopping-cart/config"
	"shopping-cart/constant"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer/user"
	"shopping-cart/repository"
	"shopping-cart/util"
)

type UserService interface {
	SaveOrUpdateUser(user *database.User) error
	ExchangeTokenAndGetProfile(code string) (*database.User, error)
	Login(req *user.Login) (string, error)
	RegisterUser(req *user.Register) error
	CreateUser(req *user.Request) error
	GetUserByID(id int) (*database.User, error)
	UpdateUser(id int, req *user.Update) error
	DeleteUser(id int) error
	GetUserByDisplayName(displayName string) (*database.User, error)
	RequestPasswordReset(email string) error
	ResetPassword(email, code, newPassword, displayName string) error
	SearchUsers(params util.SearchContainer, isMember bool) ([]database.ExternalUser, int64, error)
	GetByID(id int) (*database.ExternalUser, error)
}

type userService struct {
	userRepo   repository.UserRepository
	orderRepo  repository.OrderRepository
	verifyRepo repository.VerifyRepository
}

func NewUserService(user repository.UserRepository, order repository.OrderRepository, verify repository.VerifyRepository) UserService {
	return &userService{userRepo: user, orderRepo: order, verifyRepo: verify}
}

func (s *userService) SaveOrUpdateUser(user *database.User) error {
	return s.userRepo.Upsert(user)
}

func (s *userService) ExchangeTokenAndGetProfile(code string) (*database.User, error) {
	var tokenData *user.LineTokenResponse
	httpBuilder := builder.NewHttpClient[*user.LineTokenResponse]()

	err := httpBuilder.
		WithMethodPost().
		WithURL(constant.LineTokenURL).
		WithFormData("grant_type", "authorization_code").
		WithFormData("code", code).
		WithFormData("redirect_uri", config.AppConfig.LineRedirectURI).
		WithFormData("client_id", config.AppConfig.LineClientID).
		WithFormData("client_secret", config.AppConfig.LineClientSecret).
		UserHeaderFormUrlencoded().
		Build(&tokenData)
	if err != nil {
		return nil, err
	}

	if tokenData.AccessToken == "" {
		return nil, fmt.Errorf("failed to parse access token")
	}

	var profileData *user.LineProfileResponse
	profileData, err = util.ParseIDToken(tokenData.IDToken)
	if err != nil {
		return nil, err
	}

	user := builder.NewUserBuilder().
		WithLineID(profileData.UserID).
		WithDisplayName(profileData.DisplayName).
		WithEmail(profileData.Email).
		WithLineToken(tokenData.AccessToken).
		Build()

	return user, nil
}

func (s *userService) Login(req *user.Login) (string, error) {
	username, password := req.DisplayName, req.Password

	user, err := s.userRepo.FindByDisplayName(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := util.GenerateJWT(constant.AdminType)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) RegisterUser(req *user.Register) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	existed, _ := s.userRepo.FindByDisplayName(req.DisplayName)
	if existed != nil {
		return errors.New("username already exists")
	}

	user := builder.NewUserBuilder().
		WithLineID("CreatedByUserEmail").
		WithDisplayName(req.DisplayName).
		WithPasswordHash(string(hashedPassword)).
		WithEmail(req.Email).
		WithPhone(req.Phone).
		Build()

	return s.userRepo.Create(user)
}

func (s *userService) CreateUser(req *user.Request) error {
	user := builder.NewUserBuilder().
		WithLineID("CreatedByAdmin").
		WithDisplayName(req.DisplayName).
		WithEmail(req.Email).
		WithPhone(req.Phone).
		WithIsMember(req.IsMember).
		Build()

	return s.userRepo.Create(user)
}

func (s *userService) GetUserByID(id int) (*database.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) UpdateUser(id int, req *user.Update) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}

	updatedUser := builder.NewUserBuilder().
		WithLineID(req.LineID).
		WithDisplayName(req.DisplayName).
		WithEmail(req.Email).
		WithLineToken(req.LineToken).
		WithPhone(req.Phone).
		WithIsMember(req.IsMember).
		Build()

	updatedUser.ID = user.ID

	return s.userRepo.Update(updatedUser)
}

func (s *userService) DeleteUser(id int) error {
	tx := s.userRepo.BeginTransaction()

	pendingOrders, err := s.orderRepo.FindPendingOrdersByUserIDTx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	if len(pendingOrders) > 0 {
		tx.Rollback()
		return errors.New("userRepo has pending orders, cannot delete")
	}

	err = s.userRepo.SoftDeleteTx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *userService) GetUserByDisplayName(displayName string) (*database.User, error) {
	return s.userRepo.FindByDisplayName(displayName)
}

func (s *userService) RequestPasswordReset(email string) error {
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	err := s.verifyRepo.SaveVerificationCode(email, code)
	if err != nil {
		return err
	}

	err = util.SendResetCodeEmail(email, code)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) ResetPassword(email, code, newPassword, displayName string) error {
	savedCode, err := s.verifyRepo.GetVerificationCode(email)
	if err != nil || savedCode != code {
		return errors.New("invalid or expired verification code or code has already been used")
	}

	err = s.verifyRepo.MarkCodeAsUsed(email)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user, err := s.userRepo.FindByEmailAndDisplayName(email, displayName)
	if err != nil {
		return errors.New("user not found")
	}
	user.PasswordHash = string(hashedPassword)
	return s.userRepo.Update(user)
}

func (s *userService) SearchUsers(params util.SearchContainer, isMember bool) ([]database.ExternalUser, int64, error) {
	return s.userRepo.SearchUsers(params.Keyword, params.StartDate, params.EndDate, params.Offset, params.Limit, isMember)
}

func (s *userService) GetByID(id int) (*database.ExternalUser, error) {
	return s.userRepo.FindByIDAdmin(id)
}
