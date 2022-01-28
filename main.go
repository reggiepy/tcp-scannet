package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 22; i < 120; i++ {
		address := fmt.Sprintf("127.0.0.1:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("error creating connection: %d\n", i)
			continue
		}
		_ = conn.Close()
		fmt.Printf("Connect\n")
	}
}
