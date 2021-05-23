package data

func (pktSession *PacketSessionData) GetTrackInfo() TrackInfo {

	ti := TrackInfo{}
	ti.Era = pktSession.M_era
	ti.TotalLaps = pktSession.M_totalLaps
	ti.TrackId = pktSession.M_trackId
	ti.TrackLength = pktSession.M_trackLength

	return ti
}

func (pktSession *PacketSessionData) GetSessionTemp() SessionTemperature {

	st := SessionTemperature{}
	st.AirTemperature = pktSession.M_airTemperature
	st.TrackTemperature = pktSession.M_trackTemperature
	st.Weather = pktSession.M_weather

	return st
}
