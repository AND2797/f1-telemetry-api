package udpHeaders

func (mp *PacketMotionData) GetSuspensionPosition() []float32 {
	return mp.suspensionPosition[:]
}

func (mp *PacketMotionData) GetSuspensionVelocity() []float32 {
	return mp.suspensionVelocity[:]
}

func (mp *PacketMotionData) GetSuspensionAcceleration() []float32 {
	return mp.suspensionAcceleration[:]
}

func (mp *PacketMotionData) GetWheelSpeed() []float32 {
	return mp.wheelSpeed[:]
}

func (lp *PacketLapData) LastLapTime(carNum int) float32 {
	return lp.lapData[carNum].lastLapTime
}

func (lp *PacketLapData) GetCurrentLapTime(carNum int) float32 {
	return lp.lapData[carNum].currentLapTime
}

func (lp *PacketLapData) GetBestLapTime(carNum int) float32 {
	return lp.lapData[carNum].bestLapTime
}

func (lp *PacketLapData) GetSector1Time(carNum int) float32 {
	return lp.lapData[carNum].sector1Time
}

func (lp *PacketLapData) GetSector2Time(carNum int) float32 {
	return lp.lapData[carNum].sector2Time
}

func (lp *PacketLapData) GetLapDistance(carNum int) float32 {
	return lp.lapData[carNum].lapDistance
}

func (lp *PacketLapData) GetTotalDistance(carNum int) float32 {
	return lp.lapData[carNum].totalDistance
}

func (lp *PacketLapData) GetSafetyCarDelta(carNum int) float32 {
	return lp.lapData[carNum].safetyCarDelta
}

func (lp *PacketLapData) GetCarPosition(carNum int) uint8 {
	return lp.lapData[carNum].carPosition
}

func (lp *PacketLapData) GetCurrentLapNum(carNum int) uint8 {
	return lp.lapData[carNum].currentLapNum
}

func (lp *PacketLapData) GetPitStatus(carNum int) uint8 {
	return lp.lapData[carNum].pitStatus
}

func (lp *PacketLapData) GetSector(carNum int) uint8 {
	return lp.lapData[carNum].sector
}

func (lp *PacketLapData) GetLapInvalid(carNum int) uint8 {
	return lp.lapData[carNum].currentLapInvalid
}

func (lp *PacketLapData) GetPenalty(carNum int) uint8 {
	return lp.lapData[carNum].penalties
}

func (lp *PacketLapData) GetGridPosition(carNum int) uint8 {
	return lp.lapData[carNum].gridPosition
}

func (lp *PacketLapData) GetDriverStatus(carNum int) uint8 {
	return lp.lapData[carNum].driverStatus
}

func (lp *PacketLapData) GetResultStatus(carNum int) uint8 {
	return lp.lapData[carNum].resultStatus
}
