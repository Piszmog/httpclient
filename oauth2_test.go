package httpclient

import (
	"testing"
)

func TestCreateOAuth2Client(t *testing.T) {
	client, err := CreateOAuth2Client("clientId", "clientSecret", "accessTokenURI")
	if err != nil {
		t.Errorf("failed to create oauth2 client with error %v", err)
	}
	if client == nil {
		t.Error("no oauth2 client returned")
	}
}

func TestCreateOauth2Config(t *testing.T) {
	config, err := CreateOAuth2Config("clientId", "clientSecret", "accessTokenURI")
	if err != nil {
		t.Errorf("failed to create oauth2 with errpr %v", err)
	}
	if config == nil {
		t.Error("failed to create oauth2 config")
	}
}
