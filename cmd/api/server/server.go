// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// HTTP(s) Server
package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"golang.org/x/time/rate"

	"github.com/xwi88/api-server/cmd/api/configs"
)

// TODO: shall limit for each server
var limiter *rate.Limiter

// Server node server
type server struct {
	server      map[string]*http.Server
	httpsServer map[string]*http.Server
	conf        *configs.APIConfig
}

// NewServer new http server
func NewServer(cfg *configs.APIConfig) *server {
	s := make(map[string]*http.Server, 0)
	hs := make(map[string]*http.Server, 0)
	return &server{
		server:      s,
		httpsServer: hs,
		conf:        cfg,
	}
}

func generateHTTPServer(options configs.HTTPServerOptions) *http.Server {
	addr := fmt.Sprintf("%v:%v", options.Host, options.Port)
	// var defaultHandler http.Handler
	// if conf.EnableRateLimit {
	// 	limiter = rate.NewLimiter(rate.Limit(conf.Limiter), conf.Burst)
	// 	defaultHandler = limit(http.DefaultServeMux)
	// }
	c := cors.New(cors.Options{
		AllowedOrigins: []string{},
	})
	// handlers := c.Handler(defaultHandler)
	// router here
	handlers := c.Handler(http.DefaultServeMux)
	return &http.Server{
		Addr:           addr,
		Handler:        handlers,
		ReadTimeout:    options.ReadTimeout,
		WriteTimeout:   options.WriteTimeout,
		MaxHeaderBytes: options.MaxHeaderBytes,
	}
}

// GenerateHTTPServer generate http servers
func (s *server) GenerateHTTPServer() map[string]*http.Server {
	options := s.conf.Server
	if !options.DisableHTTP && len(options.HTTPServers) != 0 {
		log.Println("[api] create http server")
		for _, so := range options.HTTPServers {
			s.server[so.Name] = generateHTTPServer(so)
		}
	}
	return s.server
}

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			log.Println(http.StatusText(429))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Start start the server
func (s *server) Start() {
	s.GenerateHTTPServer()

	if len(s.server) == 0 && len(s.httpsServer) == 0 {
		log.Println("no valid server")
		os.Exit(1)
		return
	}

	// init the routers
	routers := &Routers{}
	routers.Init()

	if s.conf != nil && len(s.server) != 0 {
		for name, sr := range s.server {
			go startListenAndServe(name, sr)
		}
	}
}

func startListenAndServe(name string, server *http.Server) {
	log.Printf("start http server: %v, ", name)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
