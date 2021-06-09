package oauth

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("running start app tests")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestOauthConstants(t *testing.T) {

	assert.Equal(t, "X-Public", headerXPublic)
	assert.Equal(t, "X-Client-Id", headerXClientId)
	assert.Equal(t, "X-Caller-Id", headerXCallerId)
	assert.Equal(t, "access_token", paramAccessToken)
}

func TestIsPublicNilRequest(t *testing.T) {
	assert.True(t, IsPublic(nil))
}

func TestIsPublicNoError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}
	assert.False(t, IsPublic(&request))

	request.Header.Add("X-Public", "true")
	assert.True(t, IsPublic(&request))
}

func TestGetCallerIdNilRequest(t *testing.T) {
	//	TODO: implement test

}

func TestGetCallerInvalidCallerFormat(t *testing.T) {
	//	TODO: implement test

}

func TestGetCallerNoError(t *testing.T) {
	//	TODO: implement test

}

func TestGetClientIdNilRequest(t *testing.T) {
	//	TODO: implement test

}

func TestGetClientIdInvalidCallerFormat(t *testing.T) {
	//	TODO: implement test

}

func TestGetClientIdNoError(t *testing.T) {
	//	TODO: implement test

}

func TestAuthenticateRequestNilRequest(t *testing.T) {
	//	TODO: implement test

}

func TestAuthenticateRequestNoAccessToken(t *testing.T) {
	//	TODO: implement test

}

func TestAuthenticateRequestGetAccessTokenError(t *testing.T) {
	//	TODO: implement test

}

func TestAuthenticateRequestNoError(t *testing.T) {
	//	TODO: implement test
}

func TestGetAccessTokenInvalidRestclientResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodGet,
		URL:          "http://localhost:8080/oauth/access_token/Abc123",
		ReqBody:      ``,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	accessToken, err := getAccessToken("invalid restclient response when trying to get access token")
	assert.Nil(t, accessToken)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}
