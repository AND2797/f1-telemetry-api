package data

func (pktSession *PacketSessionData) GetTrackInfo() TrackInfo {

	ti := TrackInfo{}
	ti.Era = pktSession.Era
	ti.TotalLaps = pktSession.TotalLaps
	ti.TrackId = pktSession.TrackId
	ti.TrackLength = pktSession.TrackLength

	return ti
}

func (pktSession *PacketSessionData) GetSessionTemp() SessionTemperature {

	st := SessionTemperature{}
	st.AirTemperature = pktSession.AirTemperature
	st.TrackTemperature = pktSession.TrackTemperature
	st.Weather = pktSession.Weather

	return st
}
