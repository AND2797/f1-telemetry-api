package server

import (
	"f1-tele/src/udpHeaders"
	"net"
)

type Session2018 struct {
	HostPort     string
	PacketHeader udpHeaders.PacketHeader2018
	Server       *net.UDPConn
}

type Session2020 struct {
	HostPort     string
	PacketHeader udpHeaders.PacketHeader2020
	Server       *net.UDPConn
}
