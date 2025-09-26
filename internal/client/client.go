package client

import (
	"client/internal/config"
	"errors"
	"net/http"
	"time"
)

var (
	conf               = config.NewConfig()
	domain             = conf.Domain
	userEndpointsData  = conf.Routes.User
	tasksEndpointsData = conf.Routes.Task
)

// Ошибки
var (
	ErrDoRequest      = errors.New("не удалось выполнить запрос к серверу, повторите попытку позже")
	ErrDecodeResponse = errors.New("команда успешно выполнена, но получен непредвиденный ответ от сервера")
)

type AppClient struct {
	*http.Client
	Token string
}

func NewClient() *AppClient {
	return &AppClient{Client: &http.Client{Timeout: 5 * time.Second}}
}

func (c *AppClient) TokenExist() bool {
	return c.Token != ""
}
