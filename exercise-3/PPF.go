package main

import (
	"encoding/binary"
	. "fmt"
	"log"
	"net"
	"os/exec"
	"time"
)

var (
	counter uint64
	buf     = make([]byte, 16)
)

func spawnBackup() {
	(exec.Command("gnome-terminal", "--", "go", "run", "PPF.go")).Run()
}

func main() {
	addr, _ := net.ResolveUDPAddr("udp", "10.100.23.164:20007") // This needs to be fixed which each computer's id (ifconfig command)
	isPrimary := false
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Println("Error, something went wrong...")
	}
	Println("Backup Phase")
	// backup loop
	for !(isPrimary) {
		conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			isPrimary = true
		} else {
			counter = binary.BigEndian.Uint64(buf[:n])
		}
	}
	conn.Close()

	spawnBackup()
	Println("Primary Phase")
	bcastConn, _ := net.DialUDP("udp", nil, addr)
	Println("resuming from", counter)
	// primary loop
	for {
		Println("\t|", counter+1, "|")
		counter++
		binary.BigEndian.PutUint64(buf, counter)
		_, _ = bcastConn.Write(buf)
		time.Sleep(1000 * time.Millisecond)
	}
}
