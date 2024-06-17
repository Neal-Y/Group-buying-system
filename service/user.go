package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"shopping-cart/config"
	"shopping-cart/constant"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer"
	"shopping-cart/repository"
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
	resp, err := http.PostForm(constant.LineTokenURL, url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {config.AppConfig.LineRedirectURI},
		"client_id":     {config.AppConfig.LineClientID},
		"client_secret": {config.AppConfig.LineClientSecret},
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tokenData datatransfer.LineTokenResponse
	err = json.Unmarshal(body, &tokenData)
	if err != nil {
		return nil, err
	}

	if tokenData.AccessToken == "" {
		return nil, fmt.Errorf("failed to parse access token")
	}

	req, err := http.NewRequest("GET", constant.LineProfileURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenData.AccessToken))

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var profileData datatransfer.LineProfileResponse
	err = json.Unmarshal(body, &profileData)
	if err != nil {
		return nil, err
	}

	user := &database.User{
		LineID:      profileData.UserID,
		DisplayName: profileData.DisplayName,
		Email:       profileData.Email,
		LineToken:   tokenData.AccessToken,
	}

	return user, nil
}
