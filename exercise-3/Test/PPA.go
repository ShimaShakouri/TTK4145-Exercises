package main

import (
	. "fmt"
	"net"
)

func receiver() {

	ServerAddr, err := net.ResolveUDPAddr("udp", ":20007")
	if err != nil {
		Println("ResolveUDPAddr Faced Error ")
	}
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		Println("ListenUDP Faced Error ")
	}
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			Println("ReadFromUDP Faced Error ")
		}
		Println("Received ", string(buf[0:n]), " from ", addr)

	}

}

func main() {
	//go terminate()
	go receiver()
	select {}
}
