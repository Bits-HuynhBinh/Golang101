package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	/*if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}*/

	/*service := os.Args[1]

	fmt.Println(service);*/

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:12000")
	checkError(err, "ResolveTCPAddr")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, "DialTCP")

	/*_, err = conn.Write([]byte("abc"))
	checkError(err, "Write")*/

	/*result, err := ioutil.ReadAll(conn)
	checkError(err, "ReadAll")*/

	go writeData(conn);

	go readData(conn);

	//fmt.Println(string(result))

	//os.Exit(0)

	time.Sleep(time.Second * 20);
}

func writeData(conn net.Conn)  {

	for
	{
		time.Sleep(time.Second * 1);

		_, err := conn.Write([]byte("abc"))
		checkError(err, "Write")
	}

}

func readData(conn net.Conn)  {

	data := make([]byte, 128)

	for
	{
		len, err := conn.Read(data);

		if err != nil {
			fmt.Println(err)
			break
		}

		if len == 0 {
			break
		} else {
			fmt.Println(string(data[:len]));
		}

	}
}


func checkError(err error, tag string) {
	if err != nil {
		fmt.Printf(tag + "-");
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}