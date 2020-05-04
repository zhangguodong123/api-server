// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Entry for cmd, api, package name must rename to main!
package main

import (
	"github.com/xwi88/api-server/cmd/api/internal/cmd"
)

func main() {
	cmd.Execute()
}
