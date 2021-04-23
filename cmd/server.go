package main

import
(
	"fmt"
	"os"
	"net"
	"io/ioutil"
	"bytes"
	"io"
)


func main() {

	service := os.Args[1]

	// resovling address for the UDP
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClient(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr,"Error received: ", err.Error())
		os.Exit(1)
	}
}


func handleClient(conn *net.UDPConn) {

	var buf[512]byte

	_, addr, err := conn.ReadFromUDP(buf[0:])

	if err != nil {
		return
	}

	daytime := time.Now().String()

	conn.WriteToUDP([]byte(daytime), addr)


}
