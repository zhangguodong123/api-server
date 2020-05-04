// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Handler for HTTP API, helloWorld
package handler

import (
	"fmt"
	"net/http"

	"github.com/xwi88/kit4go/datetime"
)

// HelloWorldHandler helloWorld
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	now := datetime.GetNowWithZone(nil)
	resp := fmt.Sprintf("HelloWorldHandler now:%v", now.String())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}
