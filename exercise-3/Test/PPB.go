package main

import (
	. "fmt"
	"net"
	"os/exec"
	"time"
)

func main() {

	go sender()
	go jumpWindow()
	select {}
}

func sender() {

	conn, err := net.Dial("udp", "10.100.23.236:20007")
	if err != nil {
		Printf("Some error %v", err)
		return
	}
	Fprintf(conn, "%v", 5)
	time.Sleep(500 * time.Millisecond)

	for i := 0; i < 5; i++ {

		Fprintf(conn, "%v", i)
		time.Sleep(500 * time.Millisecond)
	}

	conn.Close()
}

func jumpWindow() {
	exec.Command("gnome-terminal", "--", "go", "run", "PPA.go").Run()
}
