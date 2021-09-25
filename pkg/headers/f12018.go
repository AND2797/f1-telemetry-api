package headers

// https://forums.codemasters.com/topic/30601-f1-2018-udp-specification/
type PacketHeader2018 struct {
	Header
}

type CarMotionData2018 struct {
	CarMotionData
}

type PacketMotionData2018 struct {
	PacketHeader2018
	AllCarMotionData
}

type PacketSessionData2018 struct {
	PacketHeader2018
	SessionData
}

type PacketLapData2018 struct {
	PacketHeader2018
	LapData
}

type PacketEventData2018 struct {
	PacketHeader2018
	EventData
}

type PacketParticipantData2018 struct {
	PacketHeader2018
	AllParticipantsData
}

type PacketCarSetupData2018 struct {
	PacketHeader2018
	AllCarSetupData
}

type PacketCarTelemetryData2018 struct {
	PacketHeader2018
	AllCarTelemetryData
}

type PacketCarStatusData2018 struct {
	PacketHeader2018
	AllCarStatusData
}
