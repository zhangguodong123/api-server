// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// RootCmd for api
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command called without any subCommands
var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "api command for starting the api server",
	Long:  `use "api help [<command>]" for detailed usage`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err.Error())
	}
}
