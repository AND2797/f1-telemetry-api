package headers 

type PacketHeader struct {
	M_packetFormat    uint16  // 2018
	M_packetVersion   uint8   // Version of this packet type, all start from 1
	M_packetId        uint8   // Identifier of the packet type
	M_sessionUID      uint64  // Unique identifier for the session
	M_sessionTime     float64 // session time stamp
	M_frameIdentifier uint32  // identifier for the frame the data was retrieved on
	M_playerCarIndex  uint8   // Index of player's car in the array
}

/*
Motion 0, Session 1, Lap Data 2, Event 3, Participants 4, Car Setups 5, Car Telemetry 6, Car Status 7*/

type CarMotionData struct {
	M_worldPositionX float32
	M_worldPositionY float32
	M_worldPositionZ float32

	M_worldVelocityX float32
	M_worldVelocityY float32
	M_worldVelocityZ float32

	M_worldForwardDirX int16
	M_worldForwardDirY int16
	M_worldForwardDirZ int16

	M_worldRightDirX int16
	M_worldRightDirY int16
	M_worldRightDirZ int16

	M_gForceLateral      float32
	M_gForceLongitudinal float32
	M_gForceVertical     float32
	M_yaw                float32
	M_pitch              float32
	M_roll               float32
}

type PacketMotionData struct {
	M_header        PacketHeader
	M_carMotionData [20]CarMotionData

	//Extra player car ONLY data

	M_suspensionPosition     [4]float32
	M_suspensionVelocity     [4]float32
	M_suspensionAcceleration [4]float32
	M_wheelSpeed             [4]float32
	M_wheelSlip              [4]float32
	M_localVelocityX         float32
	M_localVelocityY         float32
	M_localVelocityZ         float32
	M_angularAccelerationX   float32
	M_angularAccelerationY   float32
	M_angularAccelerationZ   float32
	M_frontWheelsAngle       float32
}

type MarshalZone struct {
	M_zoneStart float32
	M_zoneFlag  int16
}

type PacketSessionData struct {
	M_header              PacketHeader
	M_weather             uint8
	M_trackTemperature    int8
	M_airTemperature      int8
	M_totalLaps           uint8
	M_trackLength         uint16
	M_sessionType         uint8
	M_trackId             int8
	M_era                 uint8
	M_sessionTimeLeft     uint16
	M_sessionDuration     uint16
	M_PitSpeedLimit       uint8
	M_gamePaused          uint8
	M_isSpectating        uint8
	M_spectatorCarIndex   uint8
	M_sliProNativeSupport uint8
	M_marshalZones        [21]MarshalZone
	M_safetyCarStatus     uint8
	M_networkGame         uint8
}

type LapData struct {
	M_lastLapTime       float32
	M_currentLapTime    float32
	M_bestLapTime       float32
	M_sector1Time       float32
	M_sector2Time       float32
	M_lapDistance       float32
	M_totalDistance     float32
	M_SafetyCarDelta    float32
	M_carPosition       uint8
	M_currentLapNum     uint8
	M_pitStatus         uint8
	M_sector            uint8
	M_currentLapInvalid uint8
	M_penalties         uint8
	M_gridPosition      uint8
	M_driverStatus      uint8
	M_resultStatus      uint8
}

type PacketLapData struct {
	M_header  PacketHeader
	M_lapData [20]LapData
}

type PacketEventData struct {
	M_header          PacketHeader
	M_eventStringCode [4]uint8
}

type ParticipantData struct {
	M_aiControlled uint8
	M_driverId     uint8
	M_teamId       uint8
	M_raceNumber   uint8
	M_nationality  uint8
	M_name         [48]byte
}

type PacketParticipantsData struct {
	M_header      PacketHeader
	M_numCars     uint8
	M_partcipants [20]ParticipantData
}

type CarSetupData struct {
	M_frontWing             uint8
	M_rearWing              uint8
	M_onThrottle            uint8
	M_offThrottle           uint8
	M_frontCamber           float32
	M_rearCamber            float32
	M_frontToe              float32
	M_rearToe               float32
	M_frontSuspension       uint8
	M_rearSuspension        uint8
	M_frontAntiRollBar      uint8
	M_rearAntiRollBar       uint8
	M_frontSuspensionHeight uint8
	M_rearSuspensionHeight  uint8
	M_brakeBias             uint8
	M_frontTyrePressure     uint8
	M_rearTyrePressure      uint8
	M_ballast               uint8
	M_fuelLoad              float32
}

type PacketCarSetupData struct {
	M_header    PacketHeader
	M_carSetups [20]CarSetupData
}

type CarTelemetryData struct {
	M_speed                   uint16
	M_throttle                uint8
	M_steer                   int8
	M_brake                   uint8
	M_clutch                  int8
	M_gear                    int8
	M_engineRPM               uint16
	M_drs                     uint8
	M_revLightsPercent        uint8
	M_brakesTemperature       [4]uint16
	M_tyresSurfaceTemperature [4]uint16
	M_tyresInnerTemperature   [4]uint16
	M_engineTemperature       uint16
	M_tyresPressure           [4]float32
}

type PacketCarTelemetryData struct {
	M_header           PacketHeader
	M_carTelemetryData [20]CarTelemetryData
	M_buttonStatus     uint32
}

type CarStatusData struct {
	M_tractionControl         uint8
	M_antiLockBrakes          uint8
	M_fuelMix                 uint8
	M_frontBrakeBias          uint8
	M_pitLimiterStatus        uint8
	M_fuelInTank              float32
	M_fuelCapacity            float32
	M_maxRPM                  uint16
	M_idleRPM                 uint16
	M_maxGears                uint8
	M_drsAllowed              uint8
	M_tyresWear               [4]uint8
	M_tyreCompound            uint8
	M_tyresDamage             [4]uint8
	M_frontLeftWingDamage     uint8
	M_frontRightWingDamage    uint8
	M_rearWingDamage          uint8
	M_engineDamage            uint8
	M_gearBoxDamage           uint8
	M_exhaustDamage           uint8
	M_vehicleFlags            uint8
	M_ersStoreEnergy          uint8
	M_ersDeployMode           uint8
	M_ersHarvestedThisLapMGUK float32
	M_ersHarvestedThisLapMGUH float32
	M_ersDeployedThisLap      float32
}

type PacketCarStatusData struct {
	M_header        PacketHeader
	M_carStatusData [20]CarStatusData
}
