package server

import (
	"f1-tele/src/udpHeaders"
	"net"
)

type Session2018 struct {
	HostPort string
	Server   *net.UDPConn

	PacketHeader           udpHeaders.PacketHeader2018
	PacektMotionData       udpHeaders.PacketMotionData
	PacketSessionData      udpHeaders.PacketSessionData
	PacketLapData          udpHeaders.PacketLapData
	PacketEventData        udpHeaders.PacketEventData
	PacketParticipantsData udpHeaders.PacketParticipantsData
	PacketCarSetupData     udpHeaders.PacketCarSetupData
	PacketCarTelemetryData udpHeaders.PacketCarTelemetryData
	PacketCarStatusData    udpHeaders.PacketCarStatusData
}

type Session2020 struct {
	HostPort     string
	PacketHeader udpHeaders.PacketHeader2020
	Server       *net.UDPConn
}
