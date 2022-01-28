package main

import (
	"fmt"
	"log"
	"net"
	"sort"
	"time"
)

var host string
var portStart int
var portEnd int

func init() {
	host = "127.0.0.1"
	portStart = 22
	portEnd = 65535
}

func worker(ports chan int, openPorts chan []int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, time.Millisecond*300)
		if err != nil {
			//log.Printf("Failed to create connection: %d\n", port)
			openPorts <- []int{0, port}
			continue
		}
		_ = conn.Close()
		openPorts <- []int{1, port}
		//log.Printf("Connect: %d\n", port)
	}
}

func main() {
	var ports = make(chan int, 1000)
	var openPorts = make(chan []int)
	var successPort []int
	var failedPort []int
	start := time.Now()

	for i := 1; i < cap(ports); i++ {
		go worker(ports, openPorts)
	}

	go func() {
		for i := portStart; i < portEnd; i++ {
			ports <- i
		}
	}()
	//端口范围 1-65535
	//其中0不使用，1-1023为系统端口，也叫BSD保留端口
	//1024-65535为用户端口，又分为： BSD临时端口：1024-5000 BSD服务器（非特权）端口 5001-65535
	//0-1023：BSD 保留端口，也叫系统端口
	//1024-5000 BSD临时端口，一般应用程序使用1024-4999进行通信
	//5001-65535：BSD服务器（非特权）端口，用来给用户定义
	for i := portStart; i < portEnd; i++ {
		ports := <-openPorts
		status := ports[0]
		port := ports[1]
		switch status {
		case 0:
			failedPort = append(failedPort, port)
			//log.Printf("add port to failed port list: %d\n", port)
		case 1:
			successPort = append(successPort, port)
			//log.Printf("add port to success list: %d\n", port)
		default:
			//log.Fatalf("%d status not supported\n", status)
		}
	}
	close(ports)
	//log.Println("close ports")
	close(openPorts)
	//log.Println("close openPorts")

	sort.Ints(successPort)
	sort.Ints(failedPort)
	for _, port := range successPort {
		log.Printf("%s:%d success", host, port)
	}
	for _, port := range successPort {
		log.Printf("%s:%d faliled", host, port)
	}

	log.Printf("time cost: %v\n", time.Since(start))
}
