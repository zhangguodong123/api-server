// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Cmd for api
package cmd

import (
	"log"
	"net/http"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/xwi88/kit4go/datetime"

	"github.com/xwi88/api-server/cmd"
	"github.com/xwi88/api-server/cmd/api/configs"
	"github.com/xwi88/api-server/cmd/api/handler"
	"github.com/xwi88/api-server/internal/platform/utils"
)

var (
	configFile  *string
	versionFlag *bool
)

// APICmd
var APICmd = &cobra.Command{
	Use:   "start",
	Short: "start the api",
	Long: `usage example:
	api start
	start the api server`,
	Run: func(cmd *cobra.Command, args []string) {
		pid := syscall.Getpid()
		log.Printf("[api] start, pid:%v, at:%v\n", pid, datetime.GetNowWithZone(nil))

		// init config for global
		log.Println("[api] load config file:", *configFile)
		configs.Init(configFile)
		// get config
		apiConfig := configs.GetConfig()
		log.Printf("[api] Config %v", apiConfig)
		log.Println("[api] Config", configs.GetConfigJSONString(false))

		// http handler router
		http.HandleFunc("/", handler.HelloWorldHandler)
		http.HandleFunc("/hello", handler.HelloWorldHandler)
		http.HandleFunc("/ping", handler.PingHandler)

		catchSignal := utils.NewCatchSignal()
		catchSignal.RegisterSigFunc(utils.SigGroupNameBase, func() {
			log.Printf("resources close start...")
			log.Printf("resources close end...")
		}).Start()

		err := http.ListenAndServe("0.0.0.0:8080", nil)
		if err != nil {
			log.Fatalf("[api] serve err, err:%v", err.Error())
		}
	},
}

func init() {
	// add version cmd
	rootCmd.AddCommand(cmd.VersionCmd)
	versionFlag = cmd.VersionCmd.Flags().BoolP("version", "v", true, "use api version")
	cmd.VersionFlag = versionFlag

	// add api cmd
	rootCmd.AddCommand(APICmd)
	configFile = APICmd.Flags().StringP("config", "c", "./profiles/dev/api.yml", "api config file (required)")
}
