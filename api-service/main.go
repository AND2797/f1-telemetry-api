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
			switch packet := msg.(type) {
			case data.PacketCarTelemetryData2018:
				TelemetryData := packet.AllCarTelemetryData
				Car1 := TelemetryData.CarList[0]
				fmt.Println("Speed: ", Car1.Gear)
			}

		}
	}

}
