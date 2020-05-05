// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// HTTP(s) Server router
package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/xwi88/api-server/cmd/api/handler"
)

// Routers http handler set
type Routers struct {
	urls map[string]*func(http.ResponseWriter, *http.Request)
	lock sync.RWMutex
}

// Add url for routers
func (r *Routers) Add(url string, handler func(http.ResponseWriter, *http.Request)) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.urls == nil {
		r.urls = make(map[string]*func(http.ResponseWriter, *http.Request))
	}
	r.urls[url] = &handler
	http.HandleFunc(url, handler)
}

// PrintInfo print routers all
func (r *Routers) PrintInfo() {
	r.lock.Lock()
	defer r.lock.Unlock()

	var urlS string
	for k := range r.urls {
		urlS += "\n" + k
	}
	log.Printf("[api init]:routers, url: %v", urlS)
}

// Init init the routers
func (r *Routers) Init() {
	// net/http not support url regex
	r.Add("/", handler.HelloWorldHandler)
	r.Add("/api/", handler.HelloWorldHandler)
	r.Add("/api/hello", handler.HelloWorldHandler)
	r.Add("/api/ping", handler.PingHandler)

	r.PrintInfo()
}
