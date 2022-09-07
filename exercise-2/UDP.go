package main

import (
	"fmt"
	"net"
	"time"
)

func sender() {

	conn, err := net.Dial("udp", "10.100.23.240:20005")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	for {
		fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
		time.Sleep(1520 * time.Millisecond)
	}

	conn.Close()
}

func receiver() {

	ServerAddr, err := net.ResolveUDPAddr("udp", ":20005")
	if err != nil {
		fmt.Println("ResolveUDPAddr Faced Error ")
	}
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		fmt.Println("ListenUDP Faced Error ")
	}
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("ReadFromUDP Faced Error ")
		}
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

	}
}

func main() {
	go sender()
	go receiver()
	select {}
}
