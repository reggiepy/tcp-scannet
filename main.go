package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	host := "127.0.0.1"
	var wg sync.WaitGroup
	for i := 22; i < 25535; i++ {
		wg.Add(1)
		go func(port int) {
			address := fmt.Sprintf("%s:%d", host, port)
			conn, err := net.Dial("tcp", address)
			wg.Done()
			if err != nil {
				//fmt.Printf("error creating connection: %d\n", port)
				return
			}
			_ = conn.Close()
			fmt.Printf("Connect: %d\n", port)
		}(i)
	}
	wg.Wait()
}
