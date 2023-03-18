package converter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaskAuthToken(t *testing.T) {
	result := maskAuthToken("20f0c1cd-9c2a-411a-878c-9bd0bb15dc35")

	// Assert
	assert.Equal(t, "20**********************************", result)
}
