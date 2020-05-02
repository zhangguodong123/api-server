// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Version Command
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xwi88/version"
)

var (
	// VersionFlag version Flag
	VersionFlag *bool
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the bin version",
	Long: `usage example:
	<bin> version
	print the bin version`,
	Run: func(cmd *cobra.Command, args []string) {
		if *VersionFlag {
			v := version.Get()
			fmt.Println(v.StringWithIndent())
		}
	},
}
