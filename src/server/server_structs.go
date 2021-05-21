package server

import (
	"f1-tele/src/udpHeaders"
	"net"
)

type Data struct {
	PacketMotionData       udpHeaders.PacketMotionData
	PacketSessionData      udpHeaders.PacketSessionData
	PacketLapData          udpHeaders.PacketLapData
	PacketEventData        udpHeaders.PacketEventData
	PacketParticipantsData udpHeaders.PacketParticipantsData
	PacketCarSetupData     udpHeaders.PacketCarSetupData
	PacketCarTelemetryData udpHeaders.PacketCarTelemetryData
	PacketCarStatusData    udpHeaders.PacketCarStatusData

	PacketMotionChannel       chan udpHeaders.PacketMotionData
	PacketSessionChannel      chan udpHeaders.PacketSessionData
	PacketLapChannel          chan udpHeaders.PacketLapData
	PacketEventChannel        chan udpHeaders.PacketEventData
	PacketParticipantsChannel chan udpHeaders.PacketParticipantsData
	PacketCarSetupChannel     chan udpHeaders.PacketCarSetupData
	PacketCarTelemetryChannel chan udpHeaders.PacketCarTelemetryData
	PacketCarStatusChannel    chan udpHeaders.PacketCarStatusData
}

type Session2018 struct {
	HostPort     string
	Server       *net.UDPConn
	PacketHeader udpHeaders.PacketHeader2018
	Data
}

type Session2020 struct {
	HostPort     string
	PacketHeader udpHeaders.PacketHeader2020
	Server       *net.UDPConn
	Data
}
