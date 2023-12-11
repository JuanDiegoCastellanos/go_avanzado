package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	host = flag.String("host", "localhost", "host to connect")
	port = flag.Int("port", 3090, "port to connect")
)

// Client1 --> Server --> HandleConnection(Client1)
// Este handle maneja la conexion para cada cliente
func HandleConnection(conn net.Conn) {
	defer conn.Close()
	// canal por donde se envian los mensajes de ese cliente
	messagesClient := make(chan string)

	go MessageWrite(conn, messagesClient)
	// nombre del cliente tipo -> localhost:30
	clientName := conn.RemoteAddr().String()
	// Bienvenida al chat
	messagesClient <- fmt.Sprintf("Welcome to the server, your name is %s\n", clientName)
	// Notificar a todos de que se unio un nuevo cliente
	messages <- fmt.Sprintf("New client is here, name %s\n", clientName)
	// Se usa el canal de clientes entrando para enviar el canal especifico del cliente
	// que usa para enviar sus mensajes
	incomingClients <- messagesClient

	// Leer todo lo que se esta escribiendo a traves de la terminal
	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		// Envia todo el texto que se ha recolectado a traves de la terminal por el canal de mensajes del chat
		messages <- fmt.Sprintf("%s:%s\n", clientName, inputMessage.Text())
	}
	// el ciclo se rompe cuando el cliente haga control +c ctrl^C

	// Si el cliente abandona, este es enviado al canal de clientes que estan abandonando
	leavingClients <- messagesClient
	// se emite un mensaje para todo el chat, indicando que el usuario especifico dice adios y abandona
	messages <- fmt.Sprintf("%s said goodbye! ", clientName)
}

// Funcion para enviar o imprimir los mensajes atraves de la conexion
func MessageWrite(conn net.Conn, messages <-chan string) {
	for msg := range messages {
		fmt.Fprintln(conn, msg)
	}
}

func BroadCast() {
	// el total de clientes
	clients := make(map[Client]bool)
	for {
		select {
		case newMessage := <-messages:
			for client := range clients {
				// se le envia el mensaje recibido
				client <- newMessage
			}
		case newClient := <-incomingClients:
			clients[newClient] = true

		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))

	if err != nil {
		log.Fatal(err)
	}
	go BroadCast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			// fmt.Println("Has ocurred an Error")
			log.Println("Has ocurred an error: ", err)
			continue
		}
		go HandleConnection(conn)

	}
}
