package data

func (pktLap *PacketLapData) GetLapInfo(carIdx int) LapInfo {

	carIdxLap := pktLap.M_lapData[carIdx]
	lapInfo := LapInfo{}
	lapInfo.BestLapTime = carIdxLap.M_bestLapTime
	lapInfo.CurrentLapNum = carIdxLap.M_currentLapNum
	lapInfo.LastLapTime = carIdxLap.M_lastLapTime
	lapInfo.CurrentLapNum = carIdxLap.M_currentLapNum

	return lapInfo
}

func (pktLap *PacketLapData) GetSectorTimes(carIdx int) SectorTimes {

	carIdxLap := pktLap.M_lapData[carIdx]
	sectorTimes := SectorTimes{}
	sectorTimes.Sector1Time = carIdxLap.M_sector1Time
	sectorTimes.Sector2Time = carIdxLap.M_sector2Time

	return sectorTimes
}

func (pktLap *PacketLapData) GetDistances(carIdx int) Distances {

	carIdxLap := pktLap.M_lapData[carIdx]
	distances := Distances{}
	distances.LapDistance = carIdxLap.M_lapDistance
	distances.LapDistance = carIdxLap.M_totalDistance

	return distances
}
