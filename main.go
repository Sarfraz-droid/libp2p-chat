package main

import (
	"crypto/rand"
	"filelibp2p/services/connection"
	"filelibp2p/services/strio"
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("Hello, World!")

	// sourcePort := flag.Int("sp", 0, "Source port number")
	sourcePort, err := GetFreePort()

	if err != nil {
		fmt.Println("Error")
	}

	var dest string;

	fmt.Println("Enter destination port number: ")
	fmt.Scanln(&dest)


	var r io.Reader
	r = rand.Reader

	h,err := connection.MakeHost(sourcePort, r)

	if err != nil {
		fmt.Println("Error")
	}

	if dest == "" {
		connection.StartPeer(h, connection.HandleStream)
	} else {
		rw,err := connection.StartPeerAndConnect(h, dest)

		if err != nil {
			fmt.Println("Error")
		}

		go strio.ReadData(rw)
		go strio.WriteData(rw)

	}

	select {}
}

func GetFreePort() (int, error) {
        addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
        if err != nil {
                return 0, err
        }
 
        l, err := net.ListenTCP("tcp", addr)
        if err != nil {
                return 0, err
        }
        defer l.Close()
        return l.Addr().(*net.TCPAddr).Port, nil
}