// Copyright 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

// +build go1.3

package utils

import (
	"crypto/tls"
	"net/http"
	"time"
)

// NewHttpTLSTransport returns a new http.Transport constructed with the TLS config
// and the necessary parameters for Juju.
func NewHttpTLSTransport(tlsConfig *tls.Config) *http.Transport {
	// See https://code.google.com/p/go/issues/detail?id=4677
	// We need to force the connection to close each time so that we don't
	// hit the above Go bug.
	transport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		TLSClientConfig:     tlsConfig,
		DisableKeepAlives:   true,
		Dial:                dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	registerFileProtocol(transport)
	return transport
}
