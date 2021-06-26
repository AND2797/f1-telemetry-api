package server

import (
	"C"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/AND2797/f1-telemetry-api/api-service/data"
	"net"
	"strconv"
	"strings"
)

func NewSession2018(HostPort string) Session2018 {
	s2018 := Session2018{PacketsRcvd: 0, PacketsSent: 0, PacketsAcct: 0}
	s2018.DataChannel = make(chan interface{})

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
	go s2018.listen()
	return s2018
}

func decodePayload(dataStruct interface{}, payload []byte) {
	dataReader := bytes.NewReader(payload)
	err := binary.Read(dataReader, binary.LittleEndian, dataStruct)
	if err != nil {
		fmt.Println("Error in read", err)
	}
}

func (s Session2018) listen() {
	for {
		payload := make([]byte, 2048)
		_, _, _ = s.conn.ReadFromUDP(payload)
		packetHeader := data.PacketHeader2018{}
		decodePayload(&packetHeader, payload)

		switch packetID := packetHeader.PacketId; packetID {
		case data.TypeMotion:
			motionData := data.PacketMotionData2018{}
			decodePayload(&motionData, payload)
			s.DataChannel <- motionData

		case data.TypeSession:
			sessionData := data.PacketSessionData2018{}
			decodePayload(&sessionData, payload)
			s.DataChannel <- sessionData

		case data.TypeLap:
			lapData := data.PacketLapData2018{}
			decodePayload(&lapData, payload)
			s.DataChannel <- lapData

		case data.TypeEvent:
			eventData := data.PacketEventData2018{}
			decodePayload(&eventData, payload)
			s.DataChannel <- eventData

		case data.TypeParticipants:
			participantData := data.PacketParticipantData2018{}
			decodePayload(&participantData, payload)
			s.DataChannel <- participantData

		case data.TypeCarSetups:
			carSetupData := data.PacketCarSetupData2018{}
			decodePayload(&carSetupData, payload)
			s.DataChannel <- carSetupData

		case data.TypeCarTelemetry:
			carTelemetryData := data.PacketCarTelemetryData2018{}
			decodePayload(&carTelemetryData, payload)
			s.DataChannel <- carTelemetryData

		case data.TypeCarStatus:
			carStatusData := data.PacketCarStatusData2018{}
			decodePayload(&carStatusData, payload)
			s.DataChannel <- carStatusData

		default:

		}

	}
}
