package main

import (
	"encoding/base64"
	"fmt"
)

func createBasicAhthorizationHeader(userID string, authToken string) string {
	xxx := fmt.Sprintf("%s:%s", userID, authToken)
	src := []byte(xxx)
	enc := base64.StdEncoding.EncodeToString(src)
	return fmt.Sprintf("Basic %s", enc)
}
