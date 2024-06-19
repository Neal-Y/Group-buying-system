package util

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"net/http"
	"net/url"
	"shopping-cart/config"
	"shopping-cart/model/datatransfer"
)

func PostForm(url string, data url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func Get(url, accessToken string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func ParseJSONResponse(body []byte, v interface{}) error {
	err := json.Unmarshal(body, v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return nil
}

func ParseIDToken(idToken string) (*datatransfer.LineProfileResponse, error) {
	token, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.AppConfig.LineClientSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse id token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		profileData := &datatransfer.LineProfileResponse{
			UserID:      claims["sub"].(string),
			DisplayName: claims["name"].(string),
			Email:       claims["email"].(string),
		}
		return profileData, nil
	} else {
		return nil, fmt.Errorf("failed to parse id token")
	}
}
