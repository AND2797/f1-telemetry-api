package server

import (
	"bytes"
	"encoding/binary"
	"f1-tele/src/udpHeaders"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func NewSession2018(HostPort string) *Session2018 {

	s2018 := &Session2018{PacketHeader: udpHeaders.PacketHeader2018{}}
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

	return s2018
}

func (s *Session2018) Listen() {

	p := make([]byte, 2048)
	for {
		s.Server.ReadFromUDP(p)
		headerReader := bytes.NewReader(p)
		// motionReader := bytes.NewReader(p)
		// motionPacket := udpHeaders.PacketMotionData{}
		s.PacketHeader = udpHeaders.PacketHeader2018{}
		binary.Read(headerReader, binary.LittleEndian, &s.PacketHeader)
		//send info via channel?
		// switch packetId := s.PacketHeader.M_packetId; packetId{
		// case udpHeaders.Motion:
		// 	fmt.Println("motion")

		// case udpHeaders.Session:

		// case udpHeaders.Lap:

		// case udpHeaders.Event:

		// case udpHeaders.Participants:

		// case udpHeaders.CarSetups:

		// case udpHeaders.CarTelemetry:

		// case udpHeaders.

		// }
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
