package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/AND2797/f1-telemetry-api/pkg/headers"
)

func MakeChannels() DataChannels {
	dc := DataChannels{}
	dc.MotionDataChan = make(chan headers.CarMotionData)
	dc.SessionDataChan = make(chan headers.SessionData)
	dc.LapDataChan = make(chan headers.LapData)
	dc.EventDataChan = make(chan headers.EventData)
	dc.ParticipantsDataChan = make(chan headers.ParticipantData)
	dc.SetupDataChan = make(chan headers.CarSetupData)
	dc.TelemetryDataChan = make(chan headers.CarTelemetryData)
	dc.StatusDataChan = make(chan headers.CarStatusData)

	return dc
}

func NewSession2018(HostPort string) *Session2018 {
	s2018 := &Session2018{}
	s2018.DataChannels = MakeChannels()
	splitHostPort := strings.Split(HostPort, ":")
	port, _ := strconv.Atoi(splitHostPort[1])
	host := splitHostPort[0]

	addr := net.UDPAddr{Port: port, IP: net.ParseIP(host)}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println(err)
	}
	s2018.HostPort = HostPort
	s2018.conn = ser
	go s2018.Listen()
	return s2018
}

func (d *Data) decodePayload(dataStruct interface{}) {
	dataReader := bytes.NewReader(d.Payload[:])
	binary.Read(dataReader, binary.LittleEndian, dataStruct)
}

func (s Session2018) Listen() {
	for {
		payload := make([]byte, 2048)
		s.conn.ReadFromUDP(payload[:])

	}
}

// func (s *Session2018) Listen() {

// 	for {
// 		s.Server.ReadFromUDP(s.Payload[:])

// 		packetHeader := &headers.PacketHeader2018{}
// 		s.decodePayload(packetHeader)
// 		switch packetId := packetHeader.PacketId; packetId {
// 		case headers.Motion:
// 			motionPacket := headers.PacketMotionData{}
// 			s.decodePayload(&motionPacket)
// 			s.DataChannel <- &data.PacketMotionData{PacketMotionData: motionPacket}

// 		case headers.Event:
// 			eventPacket := headers.PacketEventData{}
// 			s.decodePayload(&eventPacket)
// 			s.DataChannel <- &data.PacketEventData{PacketEventData: eventPacket}

// 		case headers.Lap:
// 			lapPacket := headers.PacketLapData{}
// 			s.decodePayload(&lapPacket)
// 			s.DataChannel <- &data.PacketLapData{PacketLapData: lapPacket}

// 		case headers.CarSetups:
// 			carSetupPacket := headers.PacketCarSetupData{}
// 			s.decodePayload(&carSetupPacket)
// 			s.DataChannel <- &data.PacketCarSetupData{PacketCarSetupData: carSetupPacket}

// 		case headers.CarStatus:
// 			carStatusPacket := headers.PacketCarStatusData{}
// 			s.decodePayload(&carStatusPacket)
// 			s.DataChannel <- &data.PacketCarStatusData{PacketCarStatusData: carStatusPacket}

// 		case headers.CarTelemetry:
// 			carTelemetryPacket := headers.PacketCarTelemetryData{}
// 			s.decodePayload(&carTelemetryPacket)
// 			s.DataChannel <- &data.PacketCarTelemetryData{PacketCarTelemetryData: carTelemetryPacket}

// 		case headers.Participants:
// 			carParticipantsData := headers.ParticipantData{}
// 			s.decodePayload(&carParticipantsData)
// 			s.DataChannel <- &data.ParticipantData{ParticipantData: carParticipantsData}

// 		case headers.Session:
// 			sessionData := headers.PacketSessionData{}
// 			s.decodePayload(&sessionData)
// 			s.DataChannel <- &data.PacketSessionData{PacketSessionData: sessionData}
// 		}

// 	}
// }

func NewSession2020(HostPort string) *Session2020 {
	s2020 := &Session2020{}
	splitHostPort := strings.Split(HostPort, ":")
	port, _ := strconv.Atoi(splitHostPort[1])
	host := splitHostPort[0]

	addr := net.UDPAddr{Port: port, IP: net.ParseIP(host)}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println(err)
	}
	s2020.HostPort = HostPort
	s2020.Server = ser

	return s2020
}
