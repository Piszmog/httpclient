package net

import "testing"

func TestCreateDefaultHttpClient(t *testing.T) {
	client := CreateDefaultHttpClient()
	if client == nil {
		t.Errorf("failed to create configuration client")
	}
}
