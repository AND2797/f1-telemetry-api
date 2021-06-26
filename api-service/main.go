package main

import (
	"fmt"
	"github.com/AND2797/f1-telemetry-api/api-service/data"
	"github.com/AND2797/f1-telemetry-api/api-service/server"
	"os"
)

func main() {
	args := os.Args[1:]
	if args[0] == "2018" {
		ip := args[1]
		s2018 := server.NewSession2018(ip)
		//c2018 := client2018{server: s2018}
		for msg := range s2018.DataChannel {
			switch msg.(type) {
			case data.PacketMotionData2018:
				fmt.Println("Received PacketMotionData")

			case data.PacketSessionData2018:
				fmt.Println("Received PacketSessionData")

			case data.PacketParticipantData2018:
				fmt.Println("Received PacketParticipantData")

			case data.PacketLapData2018:
				fmt.Println("Received PacketLapData")

			case data.PacketEventData2018:
				fmt.Println("Received PacketEventData")

			case data.PacketCarSetupData2018:
				fmt.Println("Received CarSetupData")

			case data.PacketCarStatusData2018:
				fmt.Println("Received CarStatusData")

			case data.PacketCarTelemetryData2018:
				fmt.Println("Received CarTelemetryData")
			}
		}
	}
}

//type client2018 struct {
//	server server.Session2018
//}
