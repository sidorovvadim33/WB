package main

import (
	"awesomeProject/server"
	_ "github.com/lib/pq"
	"sync"
)

func init() {
	server.GetСacheFromDb()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go server.StartServer()

	server.ListenToNutsStreaming()

	wg.Wait()
}
