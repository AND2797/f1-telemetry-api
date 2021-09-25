package main

import (
	"fmt"
	"os"

	"github.com/AND2797/f1-telemetry-api/api-service/data"
	"github.com/AND2797/f1-telemetry-api/api-service/server"
)

func main() {
	args := os.Args[1:]
	if args[0] == "2018" {
		ip := args[1]
		s2018 := server.NewSession2018(ip)
		for msg := range s2018.DataChans.DataChannel {
			fmt.Println("Packets acct:", s2018.PacketsAcct)
			switch msg.(type) {
			case data.PacketMotionData2018:
				fmt.Println("YO MOTION HERE")

			}

		}
	}

}
