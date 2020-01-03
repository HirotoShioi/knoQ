package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://q.trap.jp/api/1.0"

func GetUserMe(token string) ([]byte, error) {
	return apiGetRequest(token, "/users/me")
}

func GetUsers(token string) ([]byte, error) {
	return apiGetRequest(token, "/users")
}

func apiGetRequest(token, endpoint string) ([]byte, error) {
	if token == "" {
		return nil, errors.New(http.StatusText(http.StatusUnauthorized))
	}
	req, err := http.NewRequest(http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 300 {
		return nil, errors.New(http.StatusText(res.StatusCode))
	}
	return ioutil.ReadAll(res.Body)
}
