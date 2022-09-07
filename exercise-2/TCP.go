/* GetHeadInfo
 */
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//port := ":34933"

	addr, err := net.ResolveTCPAddr("tcp", "10.100.23.240:33546")
	if err != nil {
		fmt.Println("ResolveTCPAddr Faced Error ")
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("DialTCP Faced Error ")
	}

	go writer(conn)
	go reader(conn)


	
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":17983")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	go func() {

		for {
			conn2, err := listener.AcceptTCP()
			checkError(err)

			go writer(conn2)
			go reader(conn2)
		}
	}()

	time.Sleep(150 * time.Millisecond)
	conn.Write(append([]byte("Connect to: 10.100.23.243:17983"), 0))

	//conn.Write(append([]byte("Some string"), 0))

	select {}

}

func writer(conn *net.TCPConn) {
	for {
		conn.Write(append([]byte("Some string"), 0))
		time.Sleep(1520 * time.Millisecond)
	}
}

func reader(conn *net.TCPConn) {
	for {
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		fmt.Println("message recieved:", string(buf[:n]))

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error", err)
	}
}
