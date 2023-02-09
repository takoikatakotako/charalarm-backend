package main

import (
	"encoding/base64"
	"fmt"
)
func createBasicAhthorizationHeader(userID string, authToken string) string {
	token := fmt.Sprintf("%s:%s", userID, authToken)
	src := []byte(token)
	enc := base64.StdEncoding.EncodeToString(src)
	return fmt.Sprintf("Basic %s", enc)
}
