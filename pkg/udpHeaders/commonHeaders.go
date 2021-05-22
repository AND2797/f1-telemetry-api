package udpHeaders

type PacketHeader struct {
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
	worldPositionX float32
	worldPositionY float32
	worldPositionZ float32

	worldVelocityX float32
	worldVelocityY float32
	worldVelocityZ float32

	worldForwardDirX int16
	worldForwardDirY int16
	worldForwardDirZ int16

	worldRightDirX int16
	worldRightDirY int16
	worldRightDirZ int16

	gForceLateral      float32
	gForceLongitudinal float32
	gForceVertical     float32
	yaw                float32
	pitch              float32
	roll               float32
}

type PacketMotionData struct {
	header        PacketHeader
	carMotionData [20]CarMotionData

	//Extra player car ONLY data

	suspensionPosition     [4]float32
	suspensionVelocity     [4]float32
	suspensionAcceleration [4]float32
	wheelSpeed             [4]float32
	wheelSlip              [4]float32
	localVelocityX         float32
	localVelocityY         float32
	localVelocityZ         float32
	angularAccelerationX   float32
	angularAccelerationY   float32
	angularAccelerationZ   float32
	frontWheelsAngle       float32
}

type MarshalZone struct {
	zoneStart float32
	zoneFlag  int16
}

type PacketSessionData struct {
	header              PacketHeader
	weather             uint8
	trackTemperature    int8
	airTemperature      int8
	totalLaps           uint8
	trackLength         uint16
	sessionType         uint8
	trackId             int8
	era                 uint8
	sessionTimeLeft     uint16
	sessionDuration     uint16
	pitSpeedLimit       uint8
	gamePaused          uint8
	isSpectating        uint8
	spectatorCarIndex   uint8
	sliProNativeSupport uint8
	marshalZones        [21]MarshalZone
	safetyCarStatus     uint8
	networkGame         uint8
}

type LapData struct {
	lastLapTime       float32
	currentLapTime    float32
	bestLapTime       float32
	sector1Time       float32
	sector2Time       float32
	lapDistance       float32
	totalDistance     float32
	safetyCarDelta    float32
	carPosition       uint8
	currentLapNum     uint8
	pitStatus         uint8
	sector            uint8
	currentLapInvalid uint8
	penalties         uint8
	gridPosition      uint8
	driverStatus      uint8
	resultStatus      uint8
}

type PacketLapData struct {
	Header  PacketHeader
	lapData [20]LapData
}

type PacketEventData struct {
	header          PacketHeader
	eventStringCode [4]uint8
}

type ParticipantData struct {
	aiControlled uint8
	driverId     uint8
	teamId       uint8
	raceNumber   uint8
	nationality  uint8
	name         [48]byte
}

type PacketParticipantsData struct {
	header      PacketHeader
	numCars     uint8
	partcipants [20]ParticipantData
}

type CarSetupData struct {
	frontWing             uint8
	rearWing              uint8
	onThrottle            uint8
	offThrottle           uint8
	frontCamber           float32
	rearCamber            float32
	frontToe              float32
	rearToe               float32
	frontSuspension       uint8
	rearSuspension        uint8
	frontAntiRollBar      uint8
	rearAntiRollBar       uint8
	frontSuspensionHeight uint8
	rearSuspensionHeight  uint8
	brakeBias             uint8
	frontTyrePressure     uint8
	rearTyrePressure      uint8
	ballast               uint8
	fuelLoad              float32
}

type PacketCarSetupData struct {
	header    PacketHeader
	carSetups [20]CarSetupData
}

type CarTelemetryData struct {
	speed                   uint16
	throttle                uint8
	steer                   int8
	brake                   uint8
	clutch                  int8
	gear                    int8
	engineRPM               uint16
	drs                     uint8
	revLightsPercent        uint8
	brakesTemperature       [4]uint16
	tyresSurfaceTemperature [4]uint16
	tyresInnerTemperature   [4]uint16
	engineTemperature       uint16
	tyresPressure           [4]float32
}

type PacketCarTelemetryData struct {
	header           PacketHeader
	carTelemetryData [20]CarTelemetryData
	buttonStatus     uint32
}

type CarStatusData struct {
	tractionControl         uint8
	antiLockBrakes          uint8
	fuelMix                 uint8
	frontBrakeBias          uint8
	pitLimiterStatus        uint8
	fuelInTank              float32
	fuelCapacity            float32
	maxRPM                  uint16
	idleRPM                 uint16
	maxGears                uint8
	drsAllowed              uint8
	tyresWear               [4]uint8
	tyreCompound            uint8
	tyresDamage             [4]uint8
	frontLeftWingDamage     uint8
	frontRightWingDamage    uint8
	rearWingDamage          uint8
	engineDamage            uint8
	gearBoxDamage           uint8
	exhaustDamage           uint8
	vehicleFlags            uint8
	ersStoreEnergy          uint8
	ersDeployMode           uint8
	ersHarvestedThisLapMGUK float32
	ersHarvestedThisLapMGUH float32
	ersDeployedThisLap      float32
}

type PacketCarStatusData struct {
	header        PacketHeader
	carStatusData [20]CarStatusData
}
