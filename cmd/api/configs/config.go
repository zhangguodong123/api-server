// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Configs config struct for api
package configs

import (
	"time"
)

type APIConfig struct {
	AppId  string        `json:"app_id" mapstructure:"app_id"`
	AppEnv string        `json:"app_env" mapstructure:"app_env"`
	Server ServerOptions `json:"server" mapstructure:"server"`
}

// ServerOptions http server options
type ServerOptions struct {
	DisableHTTP bool `json:"disable_http" mapstructure:"disable_http"` // default false, enable http
	EnableHTTPS bool `json:"enable_https" mapstructure:"enable_https"`
	// must use direct struct, or else not parse by yaml
	HTTPServers  []HTTPServerOptions  `json:"http" mapstructure:"http"`
	HTTPSServers []HTTPSServerOptions `json:"https" mapstructure:"https"`
}

type ServerCommonOptions struct {
	EnableRateLimit bool          `json:"enable_rate_limit" mapstructure:"enable_rate_limit"`
	Graceful        bool          `json:"graceful" mapstructure:"graceful"`
	Burst           int           `json:"burst" mapstructure:"burst"`
	MaxHeaderBytes  int           `json:"max_header_bytes" mapstructure:"max_header_bytes"`
	ReadTimeout     time.Duration `json:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `json:"write_timeout" mapstructure:"write_timeout"`
	Limiter         float64       `json:"limiter" mapstructure:"limiter"`

	Name string   `json:"name" mapstructure:"name"`
	Host string   `json:"host" mapstructure:"host"`
	Port int      `json:"port" mapstructure:"port"`
	Cors []string `json:"cors" mapstructure:"cors"`
}

// use yaml can not parse struct compose, ex: type HTTPServerOptions struct{ServerCommonOptions}
// but use the define type or alias type here can work!
// define type
// type HTTPServerOptions ServerCommonOptions
// type alias
type HTTPServerOptions = ServerCommonOptions

type HTTPSServerOptions struct {
	// ServerCommonOptions // yaml can not parse the compose struct now!!!
	EnableRateLimit bool          `json:"enable_rate_limit" mapstructure:"enable_rate_limit"`
	Graceful        bool          `json:"graceful" mapstructure:"graceful"`
	Burst           int           `json:"burst" mapstructure:"burst"`
	MaxHeaderBytes  int           `json:"max_header_bytes" mapstructure:"max_header_bytes"`
	ReadTimeout     time.Duration `json:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `json:"write_timeout" mapstructure:"write_timeout"`
	Limiter         float64       `json:"limiter" mapstructure:"limiter"`

	Name string   `json:"name" mapstructure:"name"`
	Host string   `json:"host" mapstructure:"host"`
	Port int      `json:"port" mapstructure:"port"`
	Cors []string `json:"cors" mapstructure:"cors"`

	CertFile string `json:"cert_file" mapstructure:"cert_file"`
	KeyFile  string `json:"key_file" mapstructure:"key_file"`
}
