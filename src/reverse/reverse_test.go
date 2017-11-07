package reverse

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReverseSuccessfull(t *testing.T) {
	actualResult := Reverse("Hello")
	var expectedResult = "olleH"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}



func TestReverseSuccessfull_AssertUsageExample(t *testing.T) {
	actualResult := Reverse("Hello")
	var expectedResult = "olleH"
	//t.Fatal()
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)
}