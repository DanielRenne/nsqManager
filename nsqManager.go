package main

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/DanielRenne/nsqManager/help"
)

func main() {
	c := make(chan os.Signal, 1)

	go launchNSQLookup()
	go launchNSQ()
	go launchNSQAdmin()

	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

func launchNSQLookup() {

	cmd := exec.Command(getNSQBinaryPath() + "nsqlookupd")

	for {
		cmd.Run()
	}
}

func launchNSQ() {

	cmd := exec.Command(getNSQBinaryPath()+"nsqd", "--lookupd-tcp-address=127.0.0.1:4160", "--broadcast-address=127.0.0.1")

	for {
		cmd.Run()
	}
}

func launchNSQAdmin() {

	cmd := exec.Command(getNSQBinaryPath()+"nsqadmin", "--lookupd-http-address=127.0.0.1:4161")

	for {
		cmd.Run()
	}
}

func getNSQBinaryPath() (path string) {
	path = help.GetBinaryPath() + help.PathSeparator + "bin" + help.PathSeparator + "nsq" + help.PathSeparator + "darwin" + help.PathSeparator
	return
}
