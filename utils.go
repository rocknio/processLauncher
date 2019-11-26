// +build !windows

package main

import (
	"flag"
)

var Cmd string
var H bool

func FlagsInit() {
	flag.StringVar(&Cmd, "p", "own_fog", "")
	flag.BoolVar(&H, "h", false, "this help")
}

func RunProgress() {
	childCmd := reexec.Command(Cmd)
	childCmd.Stdin = os.Stdin
	childCmd.Stdout = os.Stdout
	childCmd.Stderr = os.Stderr

	err := childCmd.Start()
	if err != nil {
		log.Panicf("failed to run command: %s", err)
	}

	err = childCmd.Wait()
	if err != nil {
		log.Panicf("failed to wait command: %s", err)
	}
}