package service

import (
	"errors"
	"fmt"
	"shopping-cart/builder"
	"shopping-cart/config"
	"shopping-cart/constant"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer/user"
	"shopping-cart/repository"
	"shopping-cart/util"
	"time"
)

type UserService interface {
	SaveOrUpdateUser(user *database.User) error
	ExchangeTokenAndGetProfile(code string) (*database.User, error)
	CreateUser(req *user.Request) error
	GetUserByID(id int) (*database.User, error)
	GetUsers() ([]database.User, error)
	UpdateUser(id int, req *user.Update) error
	DeleteUser(id int) error
}

type userService struct {
	repo  repository.UserRepository
	order repository.OrderRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) SaveOrUpdateUser(user *database.User) error {
	existingUser, err := s.repo.FindByLineID(user.LineID)
	if err != nil {
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		return s.repo.Create(user)
	}
	user.ID = existingUser.ID
	return s.repo.Update(user)
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

func (s *userService) CreateUser(req *user.Request) error {
	user := builder.NewUserBuilder().
		WithLineID("CreatedByAdmin").
		WithDisplayName(req.DisplayName).
		WithPhone(req.Phone).
		WithIsMember(req.IsMember).
		Build()

	return s.repo.Create(user)
}

func (s *userService) GetUserByID(id int) (*database.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetUsers() ([]database.User, error) {
	return s.repo.FindAll()
}

func (s *userService) UpdateUser(id int, req *user.Update) error {
	user, err := s.repo.FindByID(id)
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

	return s.repo.Update(updatedUser)
}

func (s *userService) DeleteUser(id int) error {
	tx := s.repo.BeginTransaction()

	pendingOrders, err := s.order.FindPendingOrdersByUserIDTx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	if len(pendingOrders) > 0 {
		tx.Rollback()
		return errors.New("user has pending orders, cannot delete")
	}

	err = s.repo.DeleteTx(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
