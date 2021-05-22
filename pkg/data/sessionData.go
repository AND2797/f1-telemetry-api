package data

import "github.com/AND2797/f1-telemetry-api/pkg/headers"

type TrackInfo struct {
	TotalLaps   uint8
	TrackLength uint16
	TrackId     int8
}

type SessionTemperature struct {
	Weather          uint8
	TrackTemperature int8
	AirTemperature   int8
}

type MarshalZone struct {
	headers.MarshalZone
}

type PacketSessionData struct {
	headers.PacketSessionData
}
