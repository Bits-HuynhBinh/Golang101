package main

import (
	"fmt"
	"os"
	"net"
	"time"
	"strconv"
)

func main() {
	service := "127.0.0.1:12000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError1(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError1(err)

	for {
		conn, err := listener.Accept()
		fmt.Println("Accepted");
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	defer conn.Close()

	request := make([]byte, 128)

	for
	{

		fmt.Println(1);

		len, err := conn.Read(request);

		fmt.Println(request);

		if err != nil {
			fmt.Println(err)
			break
		}

		if len == 0 {
			break // connection already closed by client
		} else if string(request[:len]) == "abc" {
			fmt.Println("true");
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
			//break
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
	}
}

func checkError1(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
