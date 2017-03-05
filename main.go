// GoTest project main.go
package main

import (
	"fmt"
	"net"
	//	"os"
	//	"time"
)

var g_bufSize int

func main() {
	fmt.Println("Hello World!")
	service := ":1201"
	g_bufSize = 512 //Set default buf size

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

func setDefaultBufSize(size int) {
	if size <= 0 {
		g_bufSize = 512
	} else {
		g_bufSize = size
	}
}
