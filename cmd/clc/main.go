package main

import (
	"fmt"
	"os"

	clc "github.com/hazelcast/hazelcast-commandline-client/clc/cmd"
)

func bye(err error) {
	_, _ = fmt.Fprintf(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	args := os.Args[1:]
	cfgPath, logPath, logLevel, err := clc.ExtractStartupArgs(args)
	if err != nil {
		bye(err)
	}
	m, err := clc.NewMain(cfgPath, logPath, logLevel)
	if err != nil {
		bye(err)
	}
	err = m.Execute(args)
	// ignoring the error here
	_ = m.Exit()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
