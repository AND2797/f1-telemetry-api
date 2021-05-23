package data

import "github.com/AND2797/f1-telemetry-api/pkg/headers"

type CarMotionData struct {
	headers.CarMotionData
}

type PacketMotionData struct {
	headers.PacketMotionData
}

type WorldPosition struct {
	WorldPositionX float32
	WorldPositionY float32
	WorldPositionZ float32
}

type WorldVelocity struct {
	WorldVelocityX float32
	WorldVelocityY float32
	WorldVelocityZ float32
}

type WorldForwardDir struct {
	WorldForwardDirX int16
	WorldForwardDirY int16
	WorldForwardDirZ int16
}

type GForces struct {
	GForceLateral      float32
	GForceLongitudinal float32
	GForceVertical     float32
}

type Attitude struct {
	Yaw   float32
	Pitch float32
	Roll  float32
}
