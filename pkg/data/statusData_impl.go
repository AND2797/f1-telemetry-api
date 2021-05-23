package data

func (pktStatus *PacketCarStatusData) GetErsStatusInfo(carIdx int) ErsStatus {
	car := pktStatus.M_carStatusData[carIdx]
	ers := ErsStatus{}
	ers.ErsDeployMode = car.M_ersDeployMode
	ers.ErsDeployedThisLap = car.M_ersDeployedThisLap
	ers.ErsHarvestedThisLapMGUH = car.M_ersHarvestedThisLapMGUH
	ers.ErsHarvestedThisLapMGUK = car.M_ersHarvestedThisLapMGUK
	ers.ErsStoreEnergy = car.M_ersStoreEnergy

	return ers
}

func (pktStatus *PacketCarStatusData) GetTyreStatus(carIdx int) TyreStatus {
	car := pktStatus.M_carStatusData[carIdx]
	ts := TyreStatus{}
	ts.TyreCompound = car.M_tyreCompound
	ts.TyreDamage = car.M_tyresDamage
	ts.TyreWear = car.M_tyresWear
	return ts
}

func (pktStatus *PacketCarStatusData) GetFuelStatus(carIdx int) FuelStatus {
	car := pktStatus.M_carStatusData[carIdx]
	fs := FuelStatus{}
	fs.FuelCapacity = car.M_fuelCapacity
	fs.FuelInTank = car.M_fuelInTank
	fs.FuelMix = car.M_fuelMix
	return fs
}

func (pktStatus *PacketCarStatusData) GetAeroStatus(carIdx int) AeroPackageStatus {
	car := pktStatus.M_carStatusData[carIdx]
	as := AeroPackageStatus{}
	as.FrontLeftWingDmg = car.M_frontLeftWingDamage
	as.FrontRightWingDmg = car.M_frontRightWingDamage
	as.RearWingDmg = car.M_rearWingDamage
	return as
}
