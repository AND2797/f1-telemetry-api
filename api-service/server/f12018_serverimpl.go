package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/AND2797/f1-telemetry-api/api-service/data"
	// "github.com/AND2797/f1-telemetry-api/pkg/headers"
)

func NewSession2018(HostPort string) *Session2018 {
	s2018 := &Session2018{PacketsRcvd: 0, PacketsSent: 0, PacketsAcct: 0}
	s2018.DataChans = Channel2018{DataChannel: make(chan interface{})}

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

func decodePayload(dataStruct interface{}, payload []byte) {
	dataReader := bytes.NewReader(payload[:])
	binary.Read(dataReader, binary.LittleEndian, dataStruct)
}

func (s *Session2018) Listen() {
	for {
		payload := make([]byte, 2048)
		s.conn.ReadFromUDP(payload[:])
		packetHeader := &data.PacketHeader2018{}
		s.PacketsRcvd += 1
		decodePayload(packetHeader, payload)
		switch packetID := packetHeader.PacketId; packetID {
		case data.TypeMotion:
			s.PacketsAcct += 1
			motionData := data.PacketMotionData2018{}
			decodePayload(&motionData, payload)
			s.DataChans.DataChannel <- motionData
			s.PacketsSent += 1

		case data.TypeSession:
			s.PacketsAcct += 1
			sessionData := data.PacketSessionData2018{}
			decodePayload(&sessionData, payload)
			s.DataChans.DataChannel <- sessionData
			s.PacketsSent += 1

		case data.TypeLap:
			s.PacketsAcct += 1
			lapData := data.PacketLapData2018{}
			decodePayload(&lapData, payload)
			s.DataChans.DataChannel <- lapData
			s.PacketsSent += 1

		case data.TypeEvent:
			s.PacketsAcct += 1
			eventData := data.PacketEventData2018{}
			decodePayload(&eventData, payload)
			s.DataChans.DataChannel <- eventData
			s.PacketsSent += 1

		case data.TypeParticipants:
			s.PacketsAcct += 1
			participantData := data.PacketParticipantData2018{}
			decodePayload(&participantData, payload)
			s.DataChans.DataChannel <- participantData
			s.PacketsSent += 1

		case data.TypeCarSetups:
			s.PacketsAcct += 1
			carSetupData := data.PacketCarSetupData2018{}
			decodePayload(&carSetupData, payload)
			s.DataChans.DataChannel <- carSetupData
			s.PacketsSent += 1

		case data.TypeCarTelemetry:
			s.PacketsAcct += 1
			carTelemetryData := data.PacketCarTelemetryData2018{}
			decodePayload(&carTelemetryData, payload)
			s.DataChans.DataChannel <- carTelemetryData
			s.PacketsSent += 1

		case data.TypeCarStatus:
			s.PacketsAcct += 1
			carStatusData := data.PacketCarStatusData2018{}
			decodePayload(&carStatusData, payload)
			s.DataChans.DataChannel <- carStatusData
			s.PacketsSent += 1

		default:

		}

	}
}
