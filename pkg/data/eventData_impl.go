package data

func (pktEvent PacketEventData) GetEventData() [4]uint8 {
	return pktEvent.EventStringCode
}
