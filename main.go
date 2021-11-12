package main

import (
	"fmt"
	"net"
)

func main() {
	s, err := net.ResolveUDPAddr("udp4", ":3333")
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(connection.LocalAddr())

	defer connection.Close()
	buffer := make([]byte, 1024)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("-> ", string(buffer[0:n-1]))

		data := []byte("yo mamma")
		fmt.Printf("data: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
