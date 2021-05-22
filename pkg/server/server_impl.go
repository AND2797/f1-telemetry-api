package server

import (
	"bytes"
	"encoding/binary"

	"github.com/AND2797/f1-telemetry-api/pkg/udpHeaders"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func NewSession2018(HostPort string) *Session2018 {
	s2018 := &Session2018{}
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
	s2018.Server = ser
	go s2018.Listen()
	return s2018
}

func (d *Data) decodePayload(dataStruct interface{}) {
	dataReader := bytes.NewReader(d.Payload[:])
	binary.Read(dataReader, binary.LittleEndian, dataStruct)
}

func (s *Session2018) Listen() {

	for {
		s.Server.ReadFromUDP(s.Payload[:])

		packetHeader := &udpHeaders.PacketHeader2018{}
		s.decodePayload(packetHeader)
		switch packetId := packetHeader.M_packetId; packetId {
		case udpHeaders.Motion:
			motionPacket := &udpHeaders.PacketMotionData{}
			s.decodePayload(motionPacket)
			s.DataChannel <- motionPacket

		case udpHeaders.Event:
			eventPacket := &udpHeaders.PacketEventData{}
			s.decodePayload(eventPacket)
			s.DataChannel <- eventPacket

		case udpHeaders.Lap:
			lapPacket := &udpHeaders.PacketLapData{}
			s.decodePayload(lapPacket)
			s.DataChannel <- lapPacket

		case udpHeaders.CarSetups:
			carSetupPacket := &udpHeaders.CarSetupData{}
			s.decodePayload(carSetupPacket)
			s.DataChannel <- carSetupPacket

		case udpHeaders.CarStatus:
			carStatusPacket := &udpHeaders.CarStatusData{}
			s.decodePayload(carStatusPacket)
			s.DataChannel <- carStatusPacket

		case udpHeaders.CarTelemetry:
			carTelemetryPacket := &udpHeaders.CarTelemetryData{}
			s.decodePayload(carTelemetryPacket)
			s.DataChannel <- carTelemetryPacket

		case udpHeaders.Participants:
			carParticipantsData := &udpHeaders.ParticipantData{}
			s.decodePayload(carParticipantsData)
			s.DataChannel <- carParticipantsData

		case udpHeaders.Session:
			sessionData := &udpHeaders.PacketSessionData{}
			s.decodePayload(sessionData)
			s.DataChannel <- sessionData
		}

	}
}

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
