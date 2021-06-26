package server

import (
	"net"
)

type Session2018 struct {
	HostPort    string
	conn        *net.UDPConn
	DataChannel chan interface{}
	PacketsRcvd int
	PacketsSent int
	PacketsAcct int
}
