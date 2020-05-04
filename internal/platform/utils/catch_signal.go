// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Common utils for catch signal after some interrupts and deal it
package utils

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	SigGroupNameBase  = "sg_base"
	SigGroupNameUsr   = "sg_usr" // syscall.SIGUSR1, syscall.SIGUSR2
	SigGroupNameOther = "sg_other"
)

var (
	sigGroupBase = [5]os.Signal{syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL}
	sigGroupUsr  = [2]os.Signal{syscall.SIGUSR1, syscall.SIGUSR2}
)

// CatchSignal signal catch struct
type CatchSignal struct {
	ch        chan os.Signal
	sigFunc   map[string]func()
	sigGroups map[string][]os.Signal
	run       bool
}

// NewCatchSignal create the catchSignal
func NewCatchSignal() *CatchSignal {
	sigGroups := make(map[string][]os.Signal)
	sigGroups[SigGroupNameBase] = sigGroupBase[:]
	sigGroups[SigGroupNameUsr] = sigGroupUsr[:]
	ch := make(chan os.Signal, 1)

	return &CatchSignal{
		sigGroups: sigGroups,
		ch:        ch,
	}
}

// RegisterSigFunc register the fun for the special signal group by sig group name
func (s *CatchSignal) RegisterSigFunc(sigGroupName string, f func()) *CatchSignal {
	if s != nil {
		if sigGroupName == SigGroupNameUsr || sigGroupName == SigGroupNameBase || sigGroupName == SigGroupNameOther {
			if len(s.sigFunc) == 0 {
				s.sigFunc = make(map[string]func())
				s.sigFunc[sigGroupName] = f
			}
		}
	}
	return s
}

// Start start catch the signal and deal it
func (s *CatchSignal) Start() {
	// SIGKILL and SIGSTOP Neither of these signals can be captured by the application,
	// nor can they be blocked or ignored by the operating system.
	// kill -9 pid => SIGKILL
	// os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2
	if s.run {
		log.Printf("CatchSignal has been started, shall not start again!")
		return
	}
	s.run = true
	signal.Notify(s.ch,
		// https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html
		syscall.SIGTERM, // "the normal way to politely ask a program to terminate"
		syscall.SIGINT,  // Ctrl+C
		syscall.SIGQUIT, // Ctrl-\
		syscall.SIGKILL, // "always fatal", "SIGKILL and SIGSTOP may not be caught by a program"
		syscall.SIGHUP,  // "terminal is disconnected"
		syscall.SIGUSR1, syscall.SIGUSR2,
	)
	go s.deal()
}

func (s *CatchSignal) deal() {
	pid := syscall.Getpid()
	for sig := range s.ch {
		switch sig {
		case syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL:
			// SigGroupNameBase
			log.Printf("kill pid:%v", pid)
			// resources close or other deals
			s.executeFunc(SigGroupNameBase)
			// kill pid real
			if err := syscall.Kill(syscall.Getpid(), syscall.SIGKILL); err != nil {
				log.Fatalf("kill pid:%v, err:%v", pid, err)
			}
			os.Exit(int(sig.(syscall.Signal)))
		case syscall.SIGUSR1, syscall.SIGUSR2:
			// SigGroupNameUsr
			log.Printf("signal:usr1|usr2 %v", sig)
			s.executeFunc(SigGroupNameUsr)
		default:
			// SigGroupNameOther
			log.Printf("signal:other %v", sig)
			s.executeFunc(SigGroupNameOther)
		}
	}
}

func (s *CatchSignal) executeFunc(sigGroupName string) {
	if len(s.sigFunc) != 0 {
		if f, exist := s.sigFunc[sigGroupName]; exist {
			f()
		}
	}
}
