package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
    "github.com/takoikatakotako/charalarm-backend/test/entity"
	"github.com/google/uuid"
)

const (
	Endpoint = "https://api.sandbox.swiswiswift.com"
)

func TestHealthCheck(t *testing.T) {
	// healthCheck
	statusCode, healthCheckResponse, err := healthcheck(t)
    if err != nil {
		t.Errorf("unexpected error: %v", err)
    }

	assert.Equal(t, statusCode, 200)
	assert.Equal(t, healthCheckResponse.Message, "healthy!")
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
	assert.Equal(t, signUpResponse.Message, "登録完了しました")


	// Withdraw
	statusCode, withdrawResponse, err := withdraw(t, userID, userToken)
    if err != nil {
		t.Errorf("unexpected error: %v", err)
    }

	assert.Equal(t, statusCode, 200)
	assert.Equal(t, withdrawResponse.Message, "退会完了しました")
}






// Get: /healthcheck
func healthcheck(t *testing.T) (int, entity.MessageResponse, error) {
	response, err := http.Get(Endpoint + "/healthcheck")
	if err != nil {
		return response.StatusCode, entity.MessageResponse{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, entity.MessageResponse{}, err
	}

	var healthCheckResponse entity.MessageResponse
	err = json.Unmarshal(body, &healthCheckResponse)
	if err != nil {
		return response.StatusCode, entity.MessageResponse{}, err
	}

	return response.StatusCode, healthCheckResponse, nil
}

// Post: /user/signup/anonymous
func signUp(t *testing.T, userID string, userToken string) (int, entity.MessageResponse, error) {
    requestBody := &entity.SignUpRequest{
        UserID: userID,
		UserToken: userToken,
    }

	jsonString, err := json.Marshal(requestBody)
    if err != nil {
		return 0, entity.MessageResponse{}, err
    }

	response, err := http.Post(Endpoint + "/user/signup/anonymous",  "application/json", bytes.NewBuffer(jsonString))
	if err != nil {
		return response.StatusCode, entity.MessageResponse{}, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, entity.MessageResponse{}, err
	}

	var signUpResponse entity.MessageResponse
	err = json.Unmarshal(responseBody, &signUpResponse)
	if err != nil {
		return response.StatusCode, entity.MessageResponse{}, err
	}

	return response.StatusCode, signUpResponse, nil
}


// POST: /user/withdraw/anonymous
func withdraw(t *testing.T, userID string, userToken string) (int, entity.MessageResponse, error) {
    requestBody := &entity.WithdrawRequest{
        UserID: userID,
		UserToken: userToken,
    }

	jsonString, err := json.Marshal(requestBody)
    if err != nil {
		return 0, entity.MessageResponse{}, err
    }

	response, err := http.Post(Endpoint + "/user/withdraw/anonymous",  "application/json", bytes.NewBuffer(jsonString))
	if err != nil {
		return response.StatusCode, entity.MessageResponse{}, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, entity.MessageResponse{}, err
	}

	var signUpResponse entity.MessageResponse
	err = json.Unmarshal(responseBody, &signUpResponse)
	if err != nil {
		return response.StatusCode, entity.MessageResponse{}, err
	}

	return response.StatusCode, signUpResponse, nil
}
