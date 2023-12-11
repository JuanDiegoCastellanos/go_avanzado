package main

import (
	"fmt"
	"net"
)

func mainx2() {
	// Escaner de puertos sin concurrencia
	for port := 0; port < 100; port++ {
		//Generar conexion tcp
		connection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", port))
		if err != nil {
			continue
		}
		connection.Close()
		fmt.Printf("Port %v is open", port)
	}
}
