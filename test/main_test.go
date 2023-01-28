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
	healthCheckResponse, err := healthcheck(t)
    if err != nil {
		t.Errorf("unexpected error: %v", err)
    }

	assert.Equal(t, healthCheckResponse.Message, "healthy!")
}

func TestSignUpAndWithdraw(t *testing.T) {
	// SingUp
	userID := uuid.New().String()
	userToken := uuid.New().String()
	signUpResponse, err := signUp(t, userID, userToken)
    if err != nil {
		t.Errorf("unexpected error: %v", err)
    }

	assert.Equal(t, signUpResponse.Message, "登録完了しました")
}



// Get: /healthcheck
func healthcheck(t *testing.T) (entity.MessageResponse, error) {
	resp, err := http.Get(Endpoint + "/healthcheck")
	if err != nil {
		return entity.MessageResponse{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.MessageResponse{}, err
	}

	var healthCheckResponse entity.MessageResponse
	err = json.Unmarshal(body, &healthCheckResponse)
	if err != nil {
		return entity.MessageResponse{}, err
	}

	assert.Equal(t, resp.StatusCode, 200)
	return healthCheckResponse, nil
}

// Post: /user/signup/anonymous
func signUp(t *testing.T, userID string, userToken string) (entity.MessageResponse, error) {
    requestBody := &entity.SignUpRequest{
        UserID: userID,
		UserToken: userToken,
    }

	jsonString, err := json.Marshal(requestBody)
    if err != nil {
		return entity.MessageResponse{}, err
    }

	response, err := http.Post(Endpoint + "/user/signup/anonymous",  "application/json", bytes.NewBuffer(jsonString))
	if err != nil {
		return entity.MessageResponse{}, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return entity.MessageResponse{}, err
	}

	var signUpResponse entity.MessageResponse
	err = json.Unmarshal(responseBody, &signUpResponse)
	if err != nil {
		return entity.MessageResponse{}, err
	}

	assert.Equal(t, response.StatusCode, 200)
	return signUpResponse, nil
}
