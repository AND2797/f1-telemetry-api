package data

import "github.com/AND2797/f1-telemetry-api/pkg/headers"

type CarSetupData struct {
	headers.CarSetupData
}

type PacketCarSetupData struct {
	headers.PacketCarSetupData
}

type AerodynamicsInfo struct {
	FrontWing uint8
	RearWing  uint8
}

type TransInfo struct {
	OnThrottle  uint8
	OffThrottle uint8
}

type SuspensionGeomInfo struct {
	FrontCamber float32
	RearCamber  float32
	FrontToe    float32
	RearToe     float32
}

type SuspensionInfo struct {
	FrontSuspension       uint8
	RearSuspension        uint8
	FrontAntiRollBar      uint8
	RearAntiRollBar       uint8
	FrontSuspensionHeight uint8
	RearSuspensionHeight  uint8
}

type BrakesInfo struct {
	BrakePressuire uint8
	BrakeBias      uint8
}

type TyresInfo struct {
	FrontTyrePressure uint8
	RearTyrePressure  uint8
}

type WeightInfo struct {
	Ballast uint8
	FuelLoad float32
}