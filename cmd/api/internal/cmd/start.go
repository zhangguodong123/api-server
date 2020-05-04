// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Cmd for api
package cmd

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/xwi88/kit4go/datetime"

	"github.com/xwi88/api-server/cmd"
	"github.com/xwi88/api-server/cmd/api/handler"
)

var (
	versionFlag *bool
)

// ApiCmd
var ApiCmd = &cobra.Command{
	Use:   "start",
	Short: "start the api",
	Long: `usage example:
	api start
	start the api server`,
	Run: func(cmd *cobra.Command, args []string) {
		pid := syscall.Getpid()
		log.Printf("[api] start, pid:%v, at:%v\n", pid, datetime.GetNowWithZone(nil))

		// http handler router
		http.HandleFunc("/", handler.HelloWorldHandler)
		http.HandleFunc("/hello", handler.HelloWorldHandler)
		http.HandleFunc("/ping", handler.PingHandler)

		ch := make(chan os.Signal, 1)
		// SIGKILL and SIGSTOP Neither of these signals can be captured by the application,
		// nor can they be blocked or ignored by the operating system.
		// kill -9 pid => SIGKILL
		// os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2
		signal.Notify(ch,
			// https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html
			syscall.SIGTERM, // "the normal way to politely ask a program to terminate"
			syscall.SIGINT,  // Ctrl+C
			syscall.SIGQUIT, // Ctrl-\
			syscall.SIGKILL, // "always fatal", "SIGKILL and SIGSTOP may not be caught by a program"
			syscall.SIGHUP,  // "terminal is disconnected"
			syscall.SIGUSR1, syscall.SIGUSR2,
		)
		go func() {
			for sig := range ch {
				switch sig {
				case syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL:
					// resources close
					log.Printf("kill pid:%v", pid)
					// kill pid real
					if err := syscall.Kill(syscall.Getpid(), syscall.SIGKILL); err != nil {
						log.Fatalf("kill pid:%v, err:%v", pid, err)
					}
					os.Exit(int(sig.(syscall.Signal)))
				case syscall.SIGUSR1:
					log.Printf("signal:usr1 %v", sig)
				case syscall.SIGUSR2:
					log.Printf("signal:usr2 %v", sig)
				default:
					log.Printf("signal:other %v", sig)
				}
			}
		}()

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
	rootCmd.AddCommand(ApiCmd)
}
