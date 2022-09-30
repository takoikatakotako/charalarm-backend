package repository

import (
	"testing"
	"fmt"
	"reflect"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreatePlatformEndpoint(t *testing.T) {
	repository := SNSRepository{IsLocal: true}

	token := uuid.New().String()
	response, err := repository.CreateIOSVoipPlatformEndpoint(token)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	response, err = repository.CreateIOSVoipPlatformEndpoint(token)
	if err != nil {

	// 	switch t := err.(type) {
	// 	default:
	// 		fmt.Println("not a model missing error")
	//    }

		fmt.Printf("type of a is %v\n", reflect.TypeOf(err))


		// fmt.Println(err.(type))
		t.Errorf("unexpected error: %v", err)
	}

	assert.NotEqual(t, len(response.EndpointArn), 100)
}
