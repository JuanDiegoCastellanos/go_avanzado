package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var site = flag.String("site", "scanme.nmap.org", "url to scan")

func mainx3() {
	// toma los parametros y los deja disponibles en la variable
	flag.Parse()
	// Escaner de puertos con concurrencia
	var wg sync.WaitGroup
	for port := 0; port < 65535; port++ {
		wg.Add(1)
		go func(portParam int) {
			defer wg.Done()
			connection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, portParam))
			if err != nil {
				return
			}
			connection.Close()
			fmt.Printf("Port %v is open \n", portParam)
		}(port)
		//Generar conexion tcp
	}
	wg.Wait()
}
