package interface_annotations

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInterfaceAnnotations(t *testing.T) {
	matryerMockBytes, err := os.ReadFile("mocks_matryer_interface_annotations_test.go")
	require.NoError(t, err)
	testifyMockBytes, err := os.ReadFile("mocks_testify_interface_annotations_test.go")
	require.NoError(t, err)
	customFileBytes, err := os.ReadFile("server_with_different_file.go")
	require.NoError(t, err)

	matryerMock := string(matryerMockBytes)
	testifyMock := string(testifyMockBytes)
	customFile := string(customFileBytes)

	assert.NotContains(t, matryerMock, "type MockRequester struct")
	assert.Contains(t, testifyMock, "type MockRequester struct")
	assert.NotContains(t, customFile, "type MockRequester struct")

	assert.NotContains(t, matryerMock, "type RequesterWithoutAnnotation struct")
	assert.NotContains(t, testifyMock, "type RequesterWithoutAnnotation struct")
	assert.NotContains(t, customFile, "type RequesterWithoutAnnotation struct")

	assert.Contains(t, matryerMock, "type MockMatryerRequester struct")
	assert.NotContains(t, testifyMock, "type MockMatryerRequester struct")
	assert.NotContains(t, customFile, "type MockMatryerRequester struct")

	assert.NotContains(t, matryerMock, "type FunServer struct")
	assert.Contains(t, testifyMock, "type FunServer struct")
	assert.NotContains(t, customFile, "type FunServer struct")

	assert.NotContains(t, matryerMock, "type MockServerWithDifferentFile struct")
	assert.NotContains(t, testifyMock, "type MockServerWithDifferentFile struct")
	assert.Contains(t, customFile, "type MockServerWithDifferentFile struct")

}
