package data

func (pktSetup *PacketCarSetupData) GetAerodynamicsInfo(carIdx int) AerodynamicsInfo {
	car := pktSetup.CarSetupsList[carIdx]
	aero := AerodynamicsInfo{}
	aero.FrontWing = car.FrontWing
	aero.RearWing = car.RearWing
	return aero
}

func (pktSetup *PacketCarSetupData) GetTransInfo(carIdx int) TransInfo {
	car := pktSetup.CarSetupsList[carIdx]
	trans := TransInfo{}
	trans.OffThrottle = car.OffThrottle
	trans.OnThrottle = car.OnThrottle
	return trans
}

func (pktSetup *PacketCarSetupData) GetSuspensionGeomInfo(carIdx int) SuspensionGeomInfo {
	car := pktSetup.CarSetupsList[carIdx]
	sg := SuspensionGeomInfo{}
	sg.FrontCamber = car.FrontCamber
	sg.FrontToe = car.FrontToe
	sg.RearCamber = car.RearCamber
	sg.RearToe = car.RearToe

	return sg
}

func (pktSetup *PacketCarSetupData) GetSuspensionInfo(carIdx int) SuspensionInfo {
	car := pktSetup.CarSetupsList[carIdx]
	si := SuspensionInfo{}
	si.FrontAntiRollBar = car.FrontAntiRollBar
	si.RearAntiRollBar = car.RearAntiRollBar
	si.FrontSuspension = car.FrontSuspension
	si.RearSuspension = car.RearSuspension
	si.FrontSuspensionHeight = car.FrontSuspensionHeight
	si.RearSuspensionHeight = car.RearSuspensionHeight
	return si
}

func (pktSetup *PacketCarSetupData) GetBrakesInfo(carIdx int) BrakesInfo {
	car := pktSetup.CarSetupsList[carIdx]
	bi := BrakesInfo{}
	bi.BrakeBias = car.BrakeBias
	bi.BrakePressuire = car.BrakePressure
	return bi
}

func (pktSetup *PacketCarSetupData) GetTyresInfo(carIdx int) TyresInfo {
	car := pktSetup.CarSetupsList[carIdx]
	ti := TyresInfo{}
	ti.FrontTyrePressure = car.FrontTyrePressure
	ti.RearTyrePressure = car.RearTyrePressure
	return ti
}

func (pktSetup *PacketCarSetupData) GetWeightInfo(carIdx int) WeightInfo {
	car := pktSetup.CarSetupsList[carIdx]
	wi := WeightInfo{}
	wi.Ballast = car.Ballast
	wi.FuelLoad = car.FuelLoad
	return wi
}
