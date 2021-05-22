package data

import "github.com/AND2797/f1-telemetry-api/pkg/headers"

type LapData struct {
	headers.LapData
}

type PacketLapData struct {
	headers.PacketLapData
}

type LapInfo struct {
	LastLapTime    float32
	CurrentLapTime float32
	BestLapTime    float32
	CurrentLapNum  uint8
}

type SectorTimes struct {
	Sector1Time float32
	Sector2Time float32
}

type Distances struct {
	LapDistance   float32
	TotalDistance float32
}
