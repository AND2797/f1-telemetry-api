package data

func (pktMotion *PacketMotionData) GetWorldPosition(carIdx int) WorldPosition {

	car := pktMotion.M_carMotionData[carIdx]
	wp := WorldPosition{}
	wp.WorldPositionX = car.M_worldPositionX
	wp.WorldPositionY = car.M_worldPositionY
	wp.WorldPositionZ = car.M_worldPositionZ

	return wp
}

func (pktMotion *PacketMotionData) GetWorldVelocity(carIdx int) WorldVelocity {

	car := pktMotion.M_carMotionData[carIdx]
	wp := WorldVelocity{}
	wp.WorldVelocityX = car.M_worldVelocityX
	wp.WorldVelocityY = car.M_worldVelocityY
	wp.WorldVelocityZ = car.M_worldVelocityZ

	return wp
}

func (pktMotion *PacketMotionData) GetWorldForwardDir(carIdx int) WorldForwardDir {

	car := pktMotion.M_carMotionData[carIdx]
	wfd := WorldForwardDir{}
	wfd.WorldForwardDirX = car.M_worldForwardDirX
	wfd.WorldForwardDirY = car.M_worldForwardDirY
	wfd.WorldForwardDirZ = car.M_worldForwardDirZ

	return wfd
}

func (pktMotion *PacketMotionData) GetGForce(carIdx int) GForces {

	car := pktMotion.M_carMotionData[carIdx]
	gf := GForces{}
	gf.GForceLateral = car.M_gForceLateral
	gf.GForceLongitudinal = car.M_gForceLongitudinal
	gf.GForceVertical = car.M_gForceVertical

	return gf
}

func (pktMotion *PacketMotionData) GetAttitude(carIdx int) Attitude {

	car := pktMotion.M_carMotionData[carIdx]
	at := Attitude{}
	at.Yaw = car.M_yaw
	at.Pitch = car.M_pitch
	at.Roll = car.M_roll

	return at
}
