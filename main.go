package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"time"
)

func checker(host string, ports, results chan int) {

	for port := range ports {
		address := net.JoinHostPort(host, strconv.Itoa(port))
		timeout := time.Second

		conn, _ := net.DialTimeout("tcp", address, timeout)

		if conn != nil {
			conn.Close()
			results <- port

			continue
		}
		results <- 0
	}
}

func main() {
	host := "localhost"
	ports := make(chan int, 100)
	results := make(chan int)
	openPorts := make([]int, 0)

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < cap(ports); i++ {
		go checker(host, ports, results)
	}

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}

	}

	close(ports)
	close(results)

	sort.Ints(openPorts)

	fmt.Println("Open ports: ", openPorts)
}
