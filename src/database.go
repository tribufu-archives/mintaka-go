// Copyright (c) Tribufu. All Rights Reserved.
// SPDX-License-Identifier: MIT

package mintaka

import (
	"errors"
	"os"
	"strconv"
)

type DatabaseDriver string

type DatabaseConfig struct {
	Driver   DatabaseDriver
	Host     string
	Port     int
	User     string
	Password string
	Schema   string
}

func NewDatabaseConfig(driver DatabaseDriver, host string, port int, user string, password string, schema string) *DatabaseConfig {
	return &DatabaseConfig{
		Driver:   driver,
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Schema:   schema,
	}
}

func NewDatabaseConfigFromEnv() (DatabaseConfig, error) {
	return NewDatabaseConfigFromEnvWithPrefix("")
}

func NewDatabaseConfigFromEnvWithPrefix(prefix string) (DatabaseConfig, error) {
	if prefix != "" {
		prefix = prefix + "_"
	}

	driverStr := os.Getenv(prefix + "DATABASE_DRIVER")
	if driverStr != "" {
		return DatabaseConfig{}, errors.New("DATABASE_DRIVER is required")
	}

	host := os.Getenv(prefix + "DATABASE_HOST")
	if host == "" {
		host = "localhost"
	}

	portStr := os.Getenv(prefix + "DATABASE_PORT")
	if portStr == "" {
		portStr = "3306"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return DatabaseConfig{}, err
	}

	return DatabaseConfig{
		Driver:   DatabaseDriver(driverStr),
		Host:     host,
		Port:     port,
		User:     os.Getenv(prefix + "DATABASE_USER"),
		Password: os.Getenv(prefix + "DATABASE_PASSWORD"),
		Schema:   os.Getenv(prefix + "DATABASE_SCHEMA"),
	}, nil
}
