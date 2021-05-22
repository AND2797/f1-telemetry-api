package data

import "github.com/AND2797/f1-telemetry-api/pkg/headers"

type CarTelemetryData struct {
	headers.CarTelemetryData
}

type PacketCarTelemetryData struct {
	headers.PacketCarTelemetryData
}
