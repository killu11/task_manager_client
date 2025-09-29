package client

import (
	"bytes"
	"client/internal/client/request"
	"client/internal/client/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var (
	createReqErr  = errors.New("failed creating request for task-api")
	failedReadErr = errors.New("failed read response body")
)

func (c *AppClient) CreateTask(
	title, description string,
	statusID int,
) (string, error) {
	reqBody, _ := json.Marshal(request.NewCreateTaskRequest(title, description, statusID))
	//-u testuser -p 1234ABc
	req, err := http.NewRequest(
		conf.Routes.Task.Create.Method,
		conf.Domain+conf.Routes.Task.Create.Path,
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return "", createReqErr
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)

	if err != nil {
		return "", fmt.Errorf("failed do request: %v", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", ErrDecodeResponse
	}

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed create task: %s", body)
	}
	return string(body), nil
}

func (c *AppClient) GetTasks() ([]*response.TaskResponse, error) {
	req, err := http.NewRequest(
		conf.Routes.Task.GetAll.Method,
		conf.Domain+conf.Routes.Task.GetAll.Path,
		nil)

	if err != nil {
		return nil, createReqErr
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", c.Token))
	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed do request: %v", err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", respBody)
	}

	taskListResponse := make([]*response.TaskResponse, 0)
	if err = json.Unmarshal(respBody, &taskListResponse); err != nil {
		return nil, fmt.Errorf("failed unmarshal response body: %v", err)
	}
	return taskListResponse, nil
}

func (c AppClient) UpdateStatusByTitle(title string, statusID int) (string, error) {
	req, err := http.NewRequest(
		conf.Routes.Task.UpdateStatus.Method,
		conf.Domain+conf.Routes.Task.UpdateStatus.Path,
		nil)

	if err != nil {
		return "", createReqErr
	}

	params := url.Values{}
	params.Set("title", title)
	params.Set("status", strconv.Itoa(statusID))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	resp, err := c.Do(req)

	if err != nil {
		return "", ErrDoRequest
	}

	defer resp.Body.Close()

	messageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed read response body")
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s", messageBytes)
	}
	return string(messageBytes), nil
}
