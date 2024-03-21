// Copyright (c) Tribufu. All Rights Reserved.
// SPDX-License-Identifier: MIT

package mintaka

import (
	"os"
	"strconv"
)

type HttpConfig struct {
	ListenIPv4           string
	EnableIPv6           bool
	ListenIPv6           string
	Port                 int
	EnableTLS            bool
	PrivateKeyFile       string
	CertificateChainFile string
	TLSPort              int
}

func NewHttpConfig(listenIPv4 string, enableIPv6 bool, listenIPv6 string, port int, enableTLS bool, privateKeyFile string, certificateChainFile string, tlsPort int) *HttpConfig {
	return &HttpConfig{
		ListenIPv4:           listenIPv4,
		EnableIPv6:           enableIPv6,
		ListenIPv6:           listenIPv6,
		Port:                 port,
		EnableTLS:            enableTLS,
		PrivateKeyFile:       privateKeyFile,
		CertificateChainFile: certificateChainFile,
		TLSPort:              tlsPort,
	}
}

func NewHttpConfigFromEnv() (HttpConfig, error) {
	return NewHttpConfigFromEnvWithPrefix("")
}

func NewHttpConfigFromEnvWithPrefix(prefix string) (HttpConfig, error) {
	if prefix != "" {
		prefix = prefix + "_"
	}

	enableIPv6Str := os.Getenv(prefix + "ENABLE_IPV6")
	if enableIPv6Str == "" {
		enableIPv6Str = "false"
	}
	enableIPv6, err := strconv.ParseBool(enableIPv6Str)
	if err != nil {
		return HttpConfig{}, err
	}

	portStr := os.Getenv(prefix + "PORT")
	if portStr == "" {
		portStr = "80"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return HttpConfig{}, err
	}

	enableTLSStr := os.Getenv(prefix + "ENABLE_TLS")
	if enableTLSStr == "" {
		enableTLSStr = "false"
	}
	enableTLS, err := strconv.ParseBool(enableTLSStr)
	if err != nil {
		return HttpConfig{}, err
	}

	tlsPortStr := os.Getenv(prefix + "TLS_PORT")
	if tlsPortStr == "" {
		tlsPortStr = "443"
	}
	tlsPort, err := strconv.Atoi(tlsPortStr)
	if err != nil {
		return HttpConfig{}, err
	}

	return HttpConfig{
		ListenIPv4:           os.Getenv(prefix + "LISTEN_IPV4"),
		EnableIPv6:           enableIPv6,
		ListenIPv6:           os.Getenv(prefix + "LISTEN_IPV6"),
		Port:                 port,
		EnableTLS:            enableTLS,
		PrivateKeyFile:       os.Getenv(prefix + "PRIVATE_KEY_FILE"),
		CertificateChainFile: os.Getenv(prefix + "CERTIFICATE_CHAIN_FILE"),
		TLSPort:              tlsPort,
	}, nil
}
