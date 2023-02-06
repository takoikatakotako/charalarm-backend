package main

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/test/entity"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	Endpoint              = "https://api.sandbox.swiswiswift.com"
	HeaderApplicationJson = "application/json"
)

func TestHealthCheck(t *testing.T) {
	// healthCheck
	statusCode, healthCheckResponse, err := healthcheck(t)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, statusCode, 200)
	assert.Equal(t, healthCheckResponse.Message, "Healthy!")
}

func TestSignUp(t *testing.T) {
	// SingUp
	userID := uuid.New().String()
	userToken := uuid.New().String()
	statusCode, signUpResponse, err := signUp(t, userID, userToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, statusCode, 200)
	assert.Equal(t, "Sign Up Success!", signUpResponse.Message)
}

func TestSignUpAndWithdraw(t *testing.T) {
	// SingUp
	userID := uuid.New().String()
	userToken := uuid.New().String()
	statusCode, signUpResponse, err := signUp(t, userID, userToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, statusCode, 200)
	assert.Equal(t, signUpResponse.Message, "Sign Up Success!")

	// Withdraw
	statusCode, withdrawResponse, err := withdraw(t, userID, userToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, statusCode, 200)
	assert.Equal(t, "Withdraw Success!", withdrawResponse.Message)
}

// Get: /healthcheck
func healthcheck(t *testing.T) (int, response.MessageResponse, error) {
	response, err := http.Get(Endpoint + "/healthcheck")
	if err != nil {
		return response.StatusCode, response.MessageResponse{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, response.MessageResponse{}, err
	}

	var healthCheckResponse response.MessageResponse
	err = json.Unmarshal(body, &healthCheckResponse)
	if err != nil {
		return response.StatusCode, response.MessageResponse{}, err
	}

	return response.StatusCode, healthCheckResponse, nil
}

// Post: /user/signup
func signUp(t *testing.T, userID string, userToken string) (int, response.MessageResponse, error) {
	requestUrl := Endpoint + "/user/signup"

	requestBody := &entity.WithdrawRequest{
		UserID:    userID,
		UserToken: userToken,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		return 0, response.MessageResponse{}, err
	}

	response, err := http.Post(requestUrl, HeaderApplicationJson, bytes.NewBuffer(jsonString))
	if err != nil {
		return response.StatusCode, response.MessageResponse{}, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, response.MessageResponse{}, err
	}

	var signUpResponse response.MessageResponse
	err = json.Unmarshal(responseBody, &signUpResponse)
	if err != nil {
		return response.StatusCode, response.MessageResponse{}, err
	}

	return response.StatusCode, signUpResponse, nil
}

// POST: /user/withdraw
func withdraw(t *testing.T, userID string, userToken string) (int, response.MessageResponse, error) {
	requestUrl := Endpoint + "/user/withdraw"

	request, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return 0, response.MessageResponse{}, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", createBasicAhthorizationHeader(userID, userToken))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return 0, response.MessageResponse{}, err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, response.MessageResponse{}, err
	}

	var signUpResponse response.MessageResponse
	err = json.Unmarshal(responseBody, &signUpResponse)
	if err != nil {
		return response.StatusCode, response.MessageResponse{}, err
	}

	return response.StatusCode, signUpResponse, nil
}

// POST: /user/info
func info(t *testing.T, userID string, userToken string) (int, response.MessageResponse, error) {
	requestBody := &entity.WithdrawRequest{
		UserID:    userID,
		UserToken: userToken,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		return 0, response.MessageResponse{}, err
	}

	response, err := http.Post(Endpoint+"/user/info", HeaderApplicationJson, bytes.NewBuffer(jsonString))
	if err != nil {
		return response.StatusCode, response.MessageResponse{}, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, response.MessageResponse{}, err
	}

	var signUpResponse response.MessageResponse
	err = json.Unmarshal(responseBody, &signUpResponse)
	if err != nil {
		return response.StatusCode, response.MessageResponse{}, err
	}

	return response.StatusCode, signUpResponse, nil
}
