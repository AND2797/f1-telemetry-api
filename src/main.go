package main

import (
	"f1-tele/src/server"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if args[0] == "2018" {
		s2018 := server.NewSession2018("127.0.0.1:20777")

		for msg := range s2018.DataChannel {
			fmt.Println(msg)
		}
	}

}
