package headers

type Header struct {
	PacketFormat    uint16  // 2018
	PacketVersion   uint8   // Version of this packet type, all start from 1
	PacketId        uint8   // Identifier of the packet type
	SessionUID      uint64  // Unique identifier for the session
	SessionTime     float64 // session time stamp
	FrameIdentifier uint32  // identifier for the frame the data was retrieved on
	PlayerCarIndex  uint8   // Index of player's car in the array
}

/*
Motion 0, Session 1, Lap Data 2, Event 3, Participants 4, Car Setups 5, Car Telemetry 6, Car Status 7*/

type CarMotionData struct {
	WorldPositionY float32
	WorldPositionX float32
	WorldPositionZ float32

	WorldVelocityX float32
	WorldVelocityY float32
	WorldVelocityZ float32

	WorldForwardDirX int16
	WorldForwardDirY int16
	WorldForwardDirZ int16

	WorldRightDirX int16
	WorldRightDirY int16
	WorldRightDirZ int16

	GForceLateral      float32
	GForceLongitudinal float32
	GForceVertical     float32
	Yaw                float32
	Pitch              float32
	Roll               float32
}

type AllCarMotionData struct {
	CarMotionDataList [20]CarMotionData

	//Extra player car ONLY data

	SuspensionPosition     [4]float32
	SuspensionVelocity     [4]float32
	SuspensionAcceleration [4]float32
	WheelSpeed             [4]float32
	WheelSlip              [4]float32
	LocalVelocityX         float32
	LocalVelocityY         float32
	LocalVelocityZ         float32
	AngularAccelerationX   float32
	AngularAccelerationY   float32
	AngularAccelerationZ   float32
	FrontWheelsAngle       float32
}

type MarshalZone struct {
	ZoneStart float32
	ZoneFlag  int16
}

type SessionData struct {
	Weather             uint8
	TrackTemperature    int8
	AirTemperature      int8
	TotalLaps           uint8
	TrackLength         uint16
	SessionType         uint8
	TrackId             int8
	Era                 uint8
	SessionTimeLeft     uint16
	SessionDuration     uint16
	PitSpeedLimit       uint8
	GamePaused          uint8
	IsSpectating        uint8
	SpectatorCarIndex   uint8
	SliProNativeSupport uint8
	MarshalZones        [21]MarshalZone
	SafetyCarStatus     uint8
	NetworkGame         uint8
}

type LapData struct {
	LastLapTime       float32
	CurrentLapTime    float32
	BestLapTime       float32
	Sector1Time       float32
	Sector2Time       float32
	LapDistance       float32
	TotalDistance     float32
	SafetyCarDelta    float32
	CarPosition       uint8
	CurrentLapNum     uint8
	PitStatus         uint8
	Sector            uint8
	CurrentLapInvalid uint8
	Penalties         uint8
	GridPosition      uint8
	DriverStatus      uint8
	ResultStatus      uint8
}

type EventData struct {
	EventStringCode [4]uint8
}

type ParticipantData struct {
	AiControlled uint8
	DriverId     uint8
	TeamId       uint8
	RaceNumber   uint8
	Nationality  uint8
	Name         [48]byte
}

type AllParticipantsData struct {
	NumCars     uint8
	Partcipants [20]ParticipantData
}

type CarSetupData struct {
	FrontWing             uint8
	RearWing              uint8
	OnThrottle            uint8
	OffThrottle           uint8
	FrontCamber           float32
	RearCamber            float32
	FrontToe              float32
	RearToe               float32
	FrontSuspension       uint8
	RearSuspension        uint8
	FrontAntiRollBar      uint8
	RearAntiRollBar       uint8
	FrontSuspensionHeight uint8
	RearSuspensionHeight  uint8
	BrakePressure         uint8
	BrakeBias             uint8
	FrontTyrePressure     uint8
	RearTyrePressure      uint8
	Ballast               uint8
	FuelLoad              float32
}

type AllCarSetupData struct {
	CarSetupsList [20]CarSetupData
}

type CarTelemetryData struct {
	Speed                   uint16
	Throttle                uint8
	Steer                   int8
	Brake                   uint8
	Clutch                  uint8
	Gear                    int8
	EngineRPM               uint16
	DRS                     uint8
	RevLightsPercent        uint8
	BrakesTemperature       [4]uint16
	TyresSurfaceTemperature [4]uint16
	TyresInnerTemperature   [4]uint16
	EngineTemperature       uint16
	TyresPressure           [4]float32
}

type AllCarTelemetryData struct {
	CarList      [20]CarTelemetryData
	ButtonStatus uint32
}

type CarStatusData struct {
	TractionControl         uint8
	AntiLockBrakes          uint8
	FuelMix                 uint8
	FrontBrakeBias          uint8
	PitLimiterStatus        uint8
	FuelInTank              float32
	FuelCapacity            float32
	MaxRPM                  uint16
	IdleRPM                 uint16
	MaxGears                uint8
	DrsAllowed              uint8
	TyresWear               [4]uint8
	TyreCompound            uint8
	TyresDamage             [4]uint8
	FrontLeftWingDamage     uint8
	FrontRightWingDamage    uint8
	RearWingDamage          uint8
	EngineDamage            uint8
	GearBoxDamage           uint8
	ExhaustDamage           uint8
	VehicleFlags            uint8
	ErsStoreEnergy          uint8
	ErsDeployMode           uint8
	ErsHarvestedThisLapMGUK float32
	ErsHarvestedThisLapMGUH float32
	ErsDeployedThisLap      float32
}

type AllCarStatusData struct {
	CarStatusDataList [20]CarStatusData
}
