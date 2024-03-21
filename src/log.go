// Copyright (c) Tribufu. All Rights Reserved.
// SPDX-License-Identifier: MIT

package mintaka

import (
	"os"
)

type LogLevel string

type LogConfig struct {
	Level LogLevel
}

func NewLogConfig(level LogLevel) *LogConfig {
	return &LogConfig{
		Level: level,
	}
}

func NewLogConfigFromEnv() (LogConfig, error) {
	return NewLogConfigFromEnvWithPrefix("")
}

func NewLogConfigFromEnvWithPrefix(prefix string) (LogConfig, error) {
	if prefix != "" {
		prefix = prefix + "_"
	}

	level := os.Getenv(prefix + "LOG_LEVEL")
	if level == "" {
		level = "info"
	}

	return LogConfig{
		Level: LogLevel(level),
	}, nil
}
