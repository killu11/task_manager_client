package client

import (
	"bytes"
	"client/internal/client/request"
	"client/internal/client/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *AppClient) SignUp(username, password string) (*response.RegisterResponse, error) {
	body, err := json.Marshal(request.NewRegisterRequest(username, password))

	if err != nil {
		log.Println("json-err: ", err)
		return nil, errors.New("не удалось отправить форму регистрации, попробуйте еще раз")
	}

	req, err := http.NewRequest( // ← добавил обработку ошибки
		userEndpointsData.SignUp.Method,
		domain+userEndpointsData.SignUp.Path,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)

	if err != nil {
		log.Println("failed post sign-up request: ", err)
		return nil, ErrDoRequest
	}
	defer resp.Body.Close()
	responseBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("не удалось создать пользователя: %s", responseBody)
	}

	regResponse := new(response.RegisterResponse)
	if json.Unmarshal(responseBody, regResponse) != nil {
		return nil, fmt.Errorf("failed decoding response body to register response: %v", err)
	}

	return regResponse, nil
}

// SignIn - send POST-request to API
func (c *AppClient) SignIn(username, password string) (*response.LoginResponse, error) {
	body := request.NewLoginRequest(username, password)
	bytesBody, err := json.Marshal(body)

	if err != nil {
		return nil, fmt.Errorf("err marshal request body to json: %v", err)
	}

	req, err := http.NewRequest(
		userEndpointsData.SignIn.Method,
		domain+userEndpointsData.SignIn.Path,
		bytes.NewBuffer(bytesBody),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed do request: %v", err)
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK || resp.Header.Get("Content-Type") == "plain/text" {
		return nil, fmt.Errorf("failed login: %stry again", respBytes)
	}

	loginResponse := new(response.LoginResponse)
	if err = json.Unmarshal(respBytes, loginResponse); err != nil {
		return nil, fmt.Errorf("failed marshal response body: %v", err)
	}
	return loginResponse, nil
}
