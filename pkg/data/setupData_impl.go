package data

func (pktSetup *PacketCarSetupData) GetAerodynamicsInfo(carIdx int) AerodynamicsInfo {
	car := pktSetup.M_carSetups[carIdx]
	aero := AerodynamicsInfo{}
	aero.FrontWing = car.M_frontWing
	aero.RearWing = car.M_rearWing
	return aero
}

func (pktSetup *PacketCarSetupData) GetTransInfo(carIdx int) TransInfo {
	car := pktSetup.M_carSetups[carIdx]
	trans := TransInfo{}
	trans.OffThrottle = car.M_offThrottle
	trans.OnThrottle = car.M_onThrottle
	return trans
}

func (pktSetup *PacketCarSetupData) GetSuspensionGeomInfo(carIdx int) SuspensionGeomInfo {
	car := pktSetup.M_carSetups[carIdx]
	sg := SuspensionGeomInfo{}
	sg.FrontCamber = car.M_frontCamber
	sg.FrontToe = car.M_frontToe
	sg.RearCamber = car.M_rearCamber
	sg.RearToe = car.M_rearToe

	return sg
}

func (pktSetup *PacketCarSetupData) GetSuspensionInfo(carIdx int) SuspensionInfo {
	car := pktSetup.M_carSetups[carIdx]
	si := SuspensionInfo{}
	si.FrontAntiRollBar = car.M_frontAntiRollBar
	si.RearAntiRollBar = car.M_rearAntiRollBar
	si.FrontSuspension = car.M_frontSuspension
	si.RearSuspension = car.M_rearSuspension
	si.FrontSuspensionHeight = car.M_frontSuspensionHeight
	si.RearSuspensionHeight = car.M_rearSuspensionHeight
	return si
}

func (pktSetup *PacketCarSetupData) GetBrakesInfo(carIdx int) BrakesInfo {
	car := pktSetup.M_carSetups[carIdx]
	bi := BrakesInfo{}
	bi.BrakeBias = car.M_brakeBias
	bi.BrakePressuire = car.M_brakePressure
	return bi
}

func (pktSetup *PacketCarSetupData) GetTyresInfo(carIdx int) TyresInfo {
	car := pktSetup.M_carSetups[carIdx]
	ti := TyresInfo{}
	ti.FrontTyrePressure = car.M_frontTyrePressure
	ti.RearTyrePressure = car.M_rearTyrePressure
	return ti
}

func (pktSetup *PacketCarSetupData) GetWeightInfo(carIdx int) WeightInfo {
	car := pktSetup.M_carSetups[carIdx]
	wi := WeightInfo{}
	wi.Ballast = car.M_ballast
	wi.FuelLoad = car.M_fuelLoad
	return wi
}
