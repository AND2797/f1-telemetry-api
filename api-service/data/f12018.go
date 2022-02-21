package data

import (
	"github.com/AND2797/f1-telemetry-api/api-service/headers"
)

// Spec sheet: https://forums.codemasters.com/topic/30601-f1-2018-udp-specification/

type PacketHeader2018 struct {
	headers.Header
}

// PacketMotionData2018 frequency: Set in Menus, Size: 1341 bytes
type PacketMotionData2018 struct {
	PacketHeader2018
	headers.AllCarMotionData
}

// PacketSessionData2018 frequency: 2 Hz, Size: 147 bytes
type PacketSessionData2018 struct {
	PacketHeader2018
	headers.SessionData
}

// PacketLapData2018 frequency: Set in Menus, Size: 841 bytes
type PacketLapData2018 struct {
	PacketHeader2018
	headers.LapData
}

// PacketEventData2018 frequency: On event, Size: 25 bytes
type PacketEventData2018 struct {
	PacketHeader2018
	headers.EventData
}

// PacketParticipantData2018 frequency:  0.2 Hz, Size: 1082 bytes
type PacketParticipantData2018 struct {
	PacketHeader2018
	headers.AllParticipantsData
}

// PacketCarSetupData2018 frequency: 2 Hz, Size: 841 bytes
type PacketCarSetupData2018 struct {
	PacketHeader2018
	headers.AllCarSetupData
}

// PacketCarTelemetryData2018 frequency: Set in Menus, Size: 1085 bytes
type PacketCarTelemetryData2018 struct {
	PacketHeader2018
	headers.AllCarTelemetryData
}

// PacketCarStatusData2018 frequency: 2 Hz, Size: 1061 bytes
type PacketCarStatusData2018 struct {
	PacketHeader2018
	headers.AllCarStatusData
}
