// GoTest project main.go
package main

import (
	"fmt"
	"net"
	//	"os"
	//	"time"
)

func main() {
	fmt.Println("Hello World!")
	service := ":1201"

	udpAddr, _ := net.ResolveUDPAddr("udp4", service)
	//checkError(err)

	conn, _ := net.ListenUDP("udp", udpAddr)
	//checkError(err)

	for {
		handleclient(conn)
	}
}

func handleclient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(addr.IP)
	fmt.Println(buf)

}
