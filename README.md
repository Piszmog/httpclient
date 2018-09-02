# Go HTTP Client
[![Build Status](https://travis-ci.org/Piszmog/httpclient.svg?branch=develop)](https://travis-ci.org/Piszmog/httpclient)
[![Coverage Status](https://coveralls.io/repos/github/Piszmog/httpclient/badge.svg?branch=develop)](https://coveralls.io/github/Piszmog/httpclient?branch=develop)
[![Go Report Card](https://goreportcard.com/badge/github.com/Piszmog/httpclient)](https://goreportcard.com/report/github.com/Piszmog/httpclient)
[![GitHub release](https://img.shields.io/github/release/Piszmog/httpclient.svg)](https://github.com/Piszmog/httpclient/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Simple module for creating HTTP clients with timeouts and connection pool.

## Description
Instead of copying and pasting the same snippet of code for creating `http.Client` with timeouts and connection pools, 
this module provides a simple package for create a client.

## Usage
There are two ways for creating a HTTP client.

### Default Client
The function `CreateDefaultHttpClient()` provides a HTTP client with timeouts set to 5 seconds, keep alive set to 30 seconds, 
TLS handshake timeout set to 5 seconds, idleConnection timeout set to 90 seconds, max idle connections is 100, and 
max idle connections per host is 100.

### Specific Client
The function `CreateHttpClient(timeout, keepAlive, tlsHandshakeTimeout, idleConnection time.Duration, idleConns, idleConnsPerHost int)` 
allows the ability to tweak the specific timeouts and connection pool.