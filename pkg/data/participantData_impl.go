package data

import "github.com/AND2797/f1-telemetry-api/pkg/headers"

func (pktPartcp *PacketParticipantsData) GetParticipantInfo(partIdx int) headers.ParticipantData {

	return pktPartcp.Partcipants[partIdx]
}
