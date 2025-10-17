package directive_comments

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInterfaceAnnotations(t *testing.T) {
	matryerMockBytes, err := os.ReadFile("mocks_matryer_directive_comments_test.go")
	require.NoError(t, err)
	testifyMockBytes, err := os.ReadFile("mocks_testify_directive_comments_test.go")
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

	assert.Contains(t, matryerMock, "type TheMatryerRequester struct")
	assert.NotContains(t, testifyMock, "type TheMatryerRequester struct")
	assert.NotContains(t, customFile, "type TheMatryerRequester struct")

	assert.NotContains(t, matryerMock, "type FunServer struct")
	assert.Contains(t, testifyMock, "type FunServer struct")
	assert.NotContains(t, customFile, "type FunServer struct")

	assert.NotContains(t, matryerMock, "type FunServerWithDifferentFile struct")
	assert.NotContains(t, testifyMock, "type FunServerWithDifferentFile struct")
	assert.Contains(t, customFile, "type FunServerWithDifferentFile struct")
	assert.NotContains(t, matryerMock, "type AnotherFunServerWithDifferentFile struct")
	assert.NotContains(t, testifyMock, "type AnotherFunServerWithDifferentFile struct")
	assert.Contains(t, customFile, "type AnotherFunServerWithDifferentFile struct")
}
