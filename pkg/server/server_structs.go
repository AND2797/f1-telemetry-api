package server

import (
	"net"

	"github.com/AND2797/f1-telemetry-api/pkg/headers"
)

type Data struct {
	PacketMotionData       headers.PacketMotionData
	PacketSessionData      headers.PacketSessionData
	PacketLapData          headers.PacketLapData
	PacketEventData        headers.PacketEventData
	PacketParticipantsData headers.PacketParticipantsData
	PacketCarSetupData     headers.PacketCarSetupData
	PacketCarTelemetryData headers.PacketCarTelemetryData
	PacketCarStatusData    headers.PacketCarStatusData
	DataChannel            chan interface{}
	Payload                [2048]byte
}

type Session2018 struct {
	HostPort string
	Server   *net.UDPConn
	Data
}

type Session2020 struct {
	HostPort string
	Server   *net.UDPConn
	Data
}
