package httpclient

import (
	"net"
	"net/http"
	"time"
)

// CreateDefaultHttpClient creates a default http.Client.
//
// Timeout set to 5 seconds, keep alive set to 30 seconds, TLS handshake timeout set to 5 seconds, idleConnection timeout set to
// 90 seconds, max idle connections is 100, and max idle connections per host is 100.
func CreateDefaultHttpClient() *http.Client {
	return CreateHttpClient(5*time.Second, 30*time.Second, 5*time.Second, 90*time.Second, 100, 100)
}

// CreateHttpClient creates a http.Client from the specified timeouts and keep alive.
//
// The client also has the maximum number of idle connections set to 100 and number of connections per host as 100.
func CreateHttpClient(timeout, keepAlive, tlsHandshakeTimeout, idleConnection time.Duration, idleConns, idleConnsPerHost int) *http.Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: keepAlive,
			DualStack: true,
		}).DialContext,
		TLSHandshakeTimeout: tlsHandshakeTimeout,
		IdleConnTimeout:     idleConnection,
		MaxIdleConns:        idleConns,
		MaxIdleConnsPerHost: idleConnsPerHost,
	}
	return &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}
}
