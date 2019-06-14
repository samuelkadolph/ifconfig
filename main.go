package main

import (
	"flag"
	"os"
)

var listens Listens

func main() {
	flag.Var(&listens, "listen", "ip:port to listen on")
	flag.Parse()

	if listens == nil {
		listens = []string{":7570"}
	}

	server, err := NewServer(listens)
	if err != nil {
		os.Exit(1)
	}

	server.Wait()
}
