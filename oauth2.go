package httpclient

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

// CreateOAuth2Client creates an OAuth2 http.Client from the provided credentials.
func CreateOAuth2Client(clientId, clientSecret, accessTokenUri string) (*http.Client, error) {
	config, err := CreateOAuth2Config(clientId, clientSecret, accessTokenUri)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create oauth2 config")
	}
	return config.Client(context.Background()), nil
}

// CreateOAuth2Config creates an OAuth2 config from the provided credentials.
func CreateOAuth2Config(clientId, clientSecret, accessTokenUri string) (*clientcredentials.Config, error) {
	return &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     accessTokenUri,
	}, nil
}
