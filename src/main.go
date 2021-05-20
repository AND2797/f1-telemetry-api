package main

import (
	"f1-tele/src/server"
	"os"
)

func main() {
	args := os.Args[1:]
	if args[0] == "2018" {
		s2018 := server.NewSession2018("127.0.0.1:20777")
		s2018.Listen()
	}

}
