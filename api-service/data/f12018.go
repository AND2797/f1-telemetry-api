package data

import (
	"github.com/AND2797/f1-telemetry-api/api-service/headers"
)

// https://forums.codemasters.com/topic/30601-f1-2018-udp-specification/
type PacketHeader2018 struct {
	headers.Header
}

type CarMotionData2018 struct {
	CarMotionData headers.CarMotionData
}

type PacketMotionData2018 struct {
	PacketHeader2018
	headers.AllCarMotionData
}

type PacketSessionData2018 struct {
	PacketHeader2018
	headers.SessionData
}

type PacketLapData2018 struct {
	PacketHeader2018
	headers.LapData
}

type PacketEventData2018 struct {
	PacketHeader2018
	headers.EventData
}

type PacketParticipantData2018 struct {
	PacketHeader2018
	headers.AllParticipantsData
}

type PacketCarSetupData2018 struct {
	PacketHeader2018
	headers.AllCarSetupData
}

type PacketCarTelemetryData2018 struct {
	PacketHeader2018
	headers.AllCarTelemetryData
}

type PacketCarStatusData2018 struct {
	PacketHeader2018
	headers.AllCarStatusData
}
