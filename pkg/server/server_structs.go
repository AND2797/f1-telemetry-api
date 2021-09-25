package server

import (
	"net"

	"github.com/AND2797/f1-telemetry-api/pkg/headers"
)

type Data struct {
	MotionData       headers.CarMotionData
	SessionData      headers.SessionData
	LapData          headers.LapData
	EventData        headers.EventData
	ParticipantsData headers.ParticipantData
	SetupData        headers.CarSetupData
	TelemetryData    headers.CarTelemetryData
	StatusData       headers.CarStatusData
}

type DataChannels struct {
	MotionDataChan       chan headers.CarMotionData
	SessionDataChan      chan headers.SessionData
	LapDataChan          chan headers.LapData
	EventDataChan        chan headers.EventData
	ParticipantsDataChan chan headers.ParticipantData
	SetupDataChan        chan headers.CarSetupData
	TelemetryDataChan    chan headers.CarTelemetryData
	StatusDataChan       chan headers.CarStatusData
}

type Session2018 struct {
	HostPort string
	conn     *net.UDPConn
	Data
	DataChannels
}

type Session2020 struct {
	HostPort string
	conn     *net.UDPConn
	Data
	DataChannels
}
