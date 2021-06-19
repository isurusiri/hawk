package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}

	conn.Close()
	return true
}

func main() {
	fmt.Println("Port Scanning")
	isOpen := scanPort("tcp", "localhost", 22)
	fmt.Printf("Port open: %t\n", isOpen)
}
