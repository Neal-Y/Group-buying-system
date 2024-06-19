package service

import (
	"fmt"
	"net/url"
	"shopping-cart/builder"
	"shopping-cart/config"
	"shopping-cart/constant"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer"
	"shopping-cart/repository"
	"shopping-cart/util"
)

type UserService interface {
	CreateUser(user *database.User) error
	GetUserByID(id int) (*database.User, error)
	UpdateUser(user *database.User) error
	DeleteUser(user *database.User) error
	FindByLineID(lineID string) (*database.User, error)
	SaveOrUpdateUser(user *database.User) error
	ExchangeTokenAndGetProfile(code string) (*database.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *database.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetUserByID(id int) (*database.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(user *database.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(user *database.User) error {
	return s.repo.Delete(user)
}

func (s *userService) FindByLineID(lineID string) (*database.User, error) {
	return s.repo.FindByLineID(lineID)
}

func (s *userService) SaveOrUpdateUser(user *database.User) error {
	existingUser, err := s.repo.FindByLineID(user.LineID)
	if err != nil {
		return s.repo.Create(user)
	}
	user.ID = existingUser.ID
	return s.repo.Update(user)
}

func (s *userService) ExchangeTokenAndGetProfile(code string) (*database.User, error) {
	body, err := util.PostForm(constant.LineTokenURL, url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {config.AppConfig.LineRedirectURI},
		"client_id":     {config.AppConfig.LineClientID},
		"client_secret": {config.AppConfig.LineClientSecret},
	})
	if err != nil {
		return nil, err
	}

	var tokenData *datatransfer.LineTokenResponse
	err = util.ParseJSONResponse(body, &tokenData)

	if tokenData.AccessToken == "" {
		return nil, fmt.Errorf("failed to parse access token")
	}

	var profileData *datatransfer.LineProfileResponse
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
