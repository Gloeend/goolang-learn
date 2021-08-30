package main

import (
	"node/api"
	"os"
	"os/signal"
	"syscall"
)

func waitForSignal() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func main() {
	go api.ServerStart()

	waitForSignal()
}
