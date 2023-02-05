package main

import (
	"fmt"
	"bytes"
	"encoding/base64"
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

// Post: /user/signup
func signUp(t *testing.T, userID string, userToken string) (int, entity.MessageResponse, error) {
	requestUrl := Endpoint + "/user/signup"

	requestBody := &entity.WithdrawRequest{
        UserID: userID,
		UserToken: userToken,
    }

	jsonString, err := json.Marshal(requestBody)
    if err != nil {
		return 0, entity.MessageResponse{}, err
    }

	response, err := http.Post(requestUrl, HeaderApplicationJson, bytes.NewBuffer(jsonString))
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

// POST: /user/withdraw
func withdraw(t *testing.T, userID string, userToken string) (int, entity.MessageResponse, error) {
	requestUrl := Endpoint + "/user/withdraw"

	request, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return 0, entity.MessageResponse{}, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", createBasicAhthorizationHeader(userID, userToken))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return 0, entity.MessageResponse{}, err
	}

	defer response.Body.Close()

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

// POST: /user/info
func info(t *testing.T, userID string, userToken string) (int, entity.MessageResponse, error) {
    requestBody := &entity.WithdrawRequest{
        UserID: userID,
		UserToken: userToken,
    }

	jsonString, err := json.Marshal(requestBody)
    if err != nil {
		return 0, entity.MessageResponse{}, err
    }

	response, err := http.Post(Endpoint + "/user/info", HeaderApplicationJson, bytes.NewBuffer(jsonString))
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



func createBasicAhthorizationHeader(userID string, authToken string) string {
	xxx := fmt.Sprintf("%s:%s", userID, authToken)
	src := []byte(xxx)
	enc := base64.StdEncoding.EncodeToString(src)
	return fmt.Sprintf("Basic %s", enc)
}

