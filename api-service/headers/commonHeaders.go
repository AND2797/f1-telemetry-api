package headers

type Header struct {
	PacketFormat    uint16  // 2018
	PacketVersion   uint8   // Version of this packet type, all start from 1
	PacketId        uint8   // Identifier of the packet type
	SessionUID      uint64  // Unique identifier for the session
	SessionTime     float32 // session time stamp
	FrameIdentifier uint32  // identifier for the frame the data was retrieved on
	PlayerCarIndex  uint8   // Index of player's car in the array
}

/*
Motion 0, Session 1, Lap Data 2, Event 3, Participants 4, Car Setups 5, Car Telemetry 6, Car Status 7*/

type CarMotionData struct {
	WorldPositionY float32 // World space X position
	WorldPositionX float32 // World space Y position
	WorldPositionZ float32 // World space Z position

	WorldVelocityX float32 // Velocity in world space X
	WorldVelocityY float32 // Velocity in world space Y
	WorldVelocityZ float32 // Velocity in world space Z

	WorldForwardDirX int16 // World space forward X direction (normalised)
	WorldForwardDirY int16 // World space forward Y direction (normalised)
	WorldForwardDirZ int16 // World space forward Z direction (normalised)

	WorldRightDirX int16 // World space right X direction (normalised)
	WorldRightDirY int16 // World space right Y direction (normalised)
	WorldRightDirZ int16 // World space right Z direction (normalised)

	GForceLateral      float32 // Lateral G-force component
	GForceLongitudinal float32 // Longitudinal G-force component
	GForceVertical     float32 // Vertical G-force component
	Yaw                float32 // Yaw angle in radians
	Pitch              float32 // Pitch angle in radians
	Roll               float32 // Roll angle in radians
}

type AllCarMotionData struct {
	CarMotionDataList [20]CarMotionData // Data for all cars on track

	//Extra player car ONLY data

	SuspensionPosition     [4]float32 // Note: All wheel arrays have the following order
	SuspensionVelocity     [4]float32 // RL, RR, FL, FR
	SuspensionAcceleration [4]float32 // RL, RR, FL, FR
	WheelSpeed             [4]float32 // Speed of each wheel
	WheelSlip              [4]float32 // Slip ratio for each wheel
	LocalVelocityX         float32    // Velocity in local space
	LocalVelocityY         float32    // Velocity in local space
	LocalVelocityZ         float32    // Velocity in local space
	AngularAccelerationX   float32    // Angular velocity x-component
	AngularAccelerationY   float32    // Angular velocity y-component
	AngularAccelerationZ   float32    // Angular velocity z-component
	FrontWheelsAngle       float32    // Current front wheels angle in radians
}

type MarshalZone struct {
	ZoneStart float32 // Fraction (0..1) of way through the lap the marshal zone starts
	ZoneFlag  int16   // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
}

type SessionData struct {
	Weather          uint8  // Weather - 0 = clear, 1 = light cloud, 2 = overcast, 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature int8   // Track temp. in degrees celsius
	AirTemperature   int8   // Air temp. in degrees celsius
	TotalLaps        uint8  // Total number of laps in this race
	TrackLength      uint16 // Track length in metres
	SessionType      uint8  // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P
	// 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ
	// 10 = R, 11 = R2, 12 = Time Trial

	TrackId             int8            // -1 for unknown, 0-21 for tracks, see appendix
	Era                 uint8           // Era, 0 = modern, 1 = classic
	SessionTimeLeft     uint16          // Time left in session in seconds
	SessionDuration     uint16          // Session duration in seconds
	PitSpeedLimit       uint8           // Pit speed limit in kilometres per hour
	GamePaused          uint8           // Whether the game is paused
	IsSpectating        uint8           // Whether the player is spectating
	SpectatorCarIndex   uint8           // Index of the car being spectated
	SliProNativeSupport uint8           // SLI Pro support, 0 = inactive, 1 = active
	NumMarshalZones     uint8           // Number of marshal zones to follow
	MarshalZones        [21]MarshalZone // List of marshal zones – max 21
	SafetyCarStatus     uint8           // 0 = no safety car, 1 = full safety car, 2 = virtual safety car
	NetworkGame         uint8           // 0 = offline, 1 = online
}

type LapData struct {
	LastLapTime    float32 // Last lap time in seconds
	CurrentLapTime float32 // Current time around the lap in seconds
	BestLapTime    float32 // Best lap time of the session in seconds
	Sector1Time    float32 // Sector 1 time in seconds
	Sector2Time    float32 // Sector 2 time in seconds
	LapDistance    float32 // Distance vehicle is around current lap in metres – could
	// be negative if line hasn’t been crossed yet

	TotalDistance float32 // Total distance travelled in session in metres – could
	// be negative if line hasn’t been crossed yet

	SafetyCarDelta    float32 // Delta in seconds for safety car
	CarPosition       uint8   // Car race position
	CurrentLapNum     uint8   // Current lap number
	PitStatus         uint8   // 0 = none, 1 = pitting, 2 = in pit area
	Sector            uint8   // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid uint8   // Current lap invalid - 0 = valid, 1 = invalid
	Penalties         uint8   // Accumulated time penalties in seconds to be added
	GridPosition      uint8   // Grid position the vehicle started the race in
	DriverStatus      uint8   // Status of driver - 0 = in garage, 1 = flying lap
	// 2 = in lap, 3 = out lap, 4 = on track

	ResultStatus uint8 // Result status - 0 = invalid, 1 = inactive, 2 = active
	// 3 = finished, 4 = disqualified, 5 = not classified
	// 6 = retired
}

type EventData struct {
	EventStringCode [4]uint8
}

type ParticipantData struct {
	AiControlled uint8    // Whether the vehicle is AI (1) or Human (0) controlled
	DriverId     uint8    // Driver id - see appendix
	TeamId       uint8    // Team id - see appendix
	RaceNumber   uint8    // Race number of the car
	Nationality  uint8    // Nationality of the driver
	Name         [48]byte // Name of participant in UTF-8 format – null terminated
	// Will be truncated with … (U+2026) if too long
}

type AllParticipantsData struct {
	NumCars      uint8
	Participants [20]ParticipantData // Number of cars in the data
}

type CarSetupData struct {
	FrontWing             uint8   // Front wing aero
	RearWing              uint8   // Rear wing aero
	OnThrottle            uint8   // Differential adjustment on throttle (percentage)
	OffThrottle           uint8   // Differential adjustment off throttle (percentage)
	FrontCamber           float32 // Front camber angle (suspension geometry)
	RearCamber            float32 // Rear camber angle (suspension geometry)
	FrontToe              float32 // Front toe angle (suspension geometry)
	RearToe               float32 // Rear toe angle (suspension geometry)
	FrontSuspension       uint8   // Front suspension
	RearSuspension        uint8   // Rear suspension
	FrontAntiRollBar      uint8   // Front anti-roll bar
	RearAntiRollBar       uint8   // Front anti-roll bar
	FrontSuspensionHeight uint8   // Front ride height
	RearSuspensionHeight  uint8   // Rear ride height
	BrakePressure         uint8   // Brake pressure (percentage)
	BrakeBias             uint8   // Brake bias (percentage)
	FrontTyrePressure     uint8   // Front tyre pressure (PSI)
	RearTyrePressure      uint8   // Rear tyre pressure (PSI)
	Ballast               uint8   // Ballast
	FuelLoad              float32 // Fuel load
}

type AllCarSetupData struct {
	CarSetupsList [20]CarSetupData
}

type CarTelemetryData struct {
	Speed                   uint16     // Speed of car in kilometres per hour
	Throttle                uint8      // Amount of throttle applied (0 to 100)
	Steer                   int8       // Steering (-100 (full lock left) to 100 (full lock right))
	Brake                   uint8      // Amount of brake applied (0 to 100)
	Clutch                  uint8      // Amount of clutch applied (0 to 100)
	Gear                    int8       // Gear selected (1-8, N=0, R=-1)
	EngineRPM               uint16     // Engine RPM
	Drs                     uint8      // 0 = off, 1 = on
	RevLightsPercent        uint8      // Rev lights indicator (percentage)
	BrakesTemperature       [4]uint16  // Brakes temperature (celsius)
	TyresSurfaceTemperature [4]uint16  // Tyres surface temperature (celsius)
	TyresInnerTemperature   [4]uint16  // Tyres inner temperature (celsius)
	EngineTemperature       uint16     // Engine temperature (celsius)
	TyresPressure           [4]float32 // Tyres pressure (PSI)
}

type AllCarTelemetryData struct {
	CarTelemetryData [20]CarTelemetryData
	ButtonStatus     uint32
}

type CarStatusData struct {
	TractionControl  uint8    // 0 (off) - 2 (high)
	AntiLockBrakes   uint8    // 0 (off) - 1 (on)
	FuelMix          uint8    // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	FrontBrakeBias   uint8    // Front brake bias (percentage)
	PitLimiterStatus uint8    // Pit limiter status - 0 = off, 1 = on
	FuelInTank       float32  // Current fuel mass
	FuelCapacity     float32  // Fuel capacity
	MaxRPM           uint16   // Cars max RPM, point of rev limiter
	IdleRPM          uint16   // Cars idle RPM
	MaxGears         uint8    // Maximum number of gears
	DrsAllowed       uint8    // 0 = not allowed, 1 = allowed, -1 = unknown
	TyresWear        [4]uint8 // Tyre wear percentage
	TyreCompound     uint8    // Modern - 0 = hyper soft, 1 = ultra soft
	// 2 = super soft, 3 = soft, 4 = medium, 5 = hard
	// 6 = super hard, 7 = inter, 8 = wet
	// Classic - 0-6 = dry, 7-8 = wet

	TyresDamage          [4]uint8 // Tyre damage (percentage)
	FrontLeftWingDamage  uint8    // Front left wing damage (percentage)
	FrontRightWingDamage uint8    // Front right wing damage (percentage)
	RearWingDamage       uint8    // Rear wing damage (percentage)
	EngineDamage         uint8    // Engine damage (percentage)
	GearBoxDamage        uint8    // Gear box damage (percentage)
	ExhaustDamage        uint8    // Exhaust damage (percentage)
	VehicleFlags         uint8    // -1 = invalid/unknown, 0 = none, 1 = green
	// 2 = blue, 3 = yellow, 4 = red

	ErsStoreEnergy uint8 // ERS energy store in Joules
	ErsDeployMode  uint8 // ERS deployment mode, 0 = none, 1 = low, 2 = medium
	// 3 = high, 4 = overtake, 5 = hotlap

	ErsHarvestedThisLapMGUK float32 // ERS energy harvested this lap by MGU-K
	ErsHarvestedThisLapMGUH float32 // ERS energy harvested this lap by MGU-H
	ErsDeployedThisLap      float32 // ERS energy deployed this lap
}

type AllCarStatusData struct {
	CarStatusDataList [20]CarStatusData
}
