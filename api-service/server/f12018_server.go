package server

import (
	"net"
)

type Session2018 struct {
	HostPort    string
	conn        *net.UDPConn
	DataChans   Channel2018
	PacketsRcvd int
	PacketsSent int
	PacketsAcct int
}

type Channel2018 struct {
	DataChannel chan interface{}
}
