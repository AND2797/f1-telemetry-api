package main

import (
	"bytes"
	"encoding/binary"
	"f1-tele/src/f12018"
	"fmt"
	"net"
)

func main() {
	p := make([]byte, 2048)
	addr := net.UDPAddr{Port: 20777, IP: net.ParseIP("127.0.0.1")}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Some error", err)
		return
	}

	for {
		_, _, err := ser.ReadFromUDP(p)
		buf := bytes.NewReader(p)
		pktHeader := f12018.PacketHeader{}
		err = binary.Read(buf, binary.LittleEndian, &pktHeader)
		// pktHeader, err = getPacketHeader(p, pktHeader)
		fmt.Println(pktHeader)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}
