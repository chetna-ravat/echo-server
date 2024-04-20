package server

import (
	"io"
	"errors"
	"strconv"
	"net"
	"log"
	"syscall"

	"github.com/chetna-ravat/echo-server/config"
)

func readMessage(c net.Conn) (string, error) {

	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func respond(msg string, c net.Conn) error {
	if _, err := c.Write([]byte(msg)); err != nil {
		return err
	}
	return nil
}

func RunSyncServer() {
	log.Println("Starting a synchronous TCP server on", config.Host, config.Port)

	var con_clients int = 0
	// start listening to the configured host and port
	listner, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}

	for {
		// Blocking call: wait for new client to connect
		connection, err := listner.Accept()
		if err != nil {
			panic(err)
		}

		// Bump up the client connection count
		con_clients += 1
		log.Println("Client connected with address: ", connection.RemoteAddr(), "Total connected client: ", con_clients)

		for {
			// over the socket, continuously read the command and print it out
			msg, err := readMessage(connection)
			if err != nil {
				// Found error when reading message from client
				connection.Close()
				con_clients -= 1
				log.Println("Client disconnected", connection.RemoteAddr(), "Total connected client: ", con_clients)
				if errors.Is(err, syscall.ECONNRESET) || errors.Is(err, io.EOF) {
					break
				}
				log.Println("err", err)
			}
			log.Println("Message: ", msg)
			if err = respond(msg, connection); err != nil {
				log.Println("Error in echoing: ", err)
			}
		}
	}
}
