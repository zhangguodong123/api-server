// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Version Command to build the bin
package main

import (
	"github.com/xwi88/api-server/cmd"
)

var (
	versionFlag *bool
)

func main() {
	versionFlag = cmd.VersionCmd.Flags().BoolP("version", "v", true, "")
	cmd.VersionFlag = versionFlag

	if err := cmd.VersionCmd.Execute(); err != nil {
		panic(err.Error())
	}
}
