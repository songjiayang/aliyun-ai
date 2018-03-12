package client

import (
	"net/http"
	"testing"
	"time"

	"github.com/golib/assert"
)

func TestNewClient(t *testing.T) {
	assertion := assert.New(t)
	authClient := New(NewConfig("", "", "", 30*time.Second))
	assertion.NotNil(authClient)
}

func TestSend(t *testing.T) {
	assertion := assert.New(t)
	authClient := New(NewConfig("", "", "", 30*time.Second))

	body, err := authClient.Send(http.MethodPost, "http://example.com", "xxx", nil)

	assertion.Nil(err)
	assertion.NotNil(body)
}
