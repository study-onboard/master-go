package server

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// run server
func Run(port int) {
	fmt.Printf("Starting server on port: %d\n", port)

	l, err := net.Listen("tcp4", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("accept connection failed: %s\n", err.Error())
			continue
		}

		go handleConnection(conn)
	}
}

// handle connection
//
// package struct
// -------------------------------------------
// uint16 - package type
// uint32 - payload size
// bytes[size] - payload
func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("handle connection from %s\n", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	for {
		// read package type
		buffer := make([]byte, 2)
		err := readBytes(reader, buffer)
		if err != nil {
			fmt.Printf(
				"read package type from connection %s failed: %s\n",
				conn.RemoteAddr().String(),
				err.Error(),
			)
			break
		}

		packageType := binary.BigEndian.Uint16(buffer)
		if packageType != 0 && packageType != 1 {
			fmt.Printf("invalid package type from connection: %s\n", conn.RemoteAddr().String())
			return
		}
		// health ping package
		if packageType == 0 {
			fmt.Printf("health ping from %s", conn.RemoteAddr().String())
			// read more extra bytes for health package
			buffer = make([]byte, 2)
			readBytes(reader, buffer)
			continue
		}

		// read payload size
		buffer = make([]byte, 4)
		err = readBytes(reader, buffer)
		if err != nil {
			fmt.Printf(
				"read payload size from connection %s failed: %s\n",
				conn.RemoteAddr().String(),
				err.Error(),
			)
			break
		}
		payloadSize := binary.BigEndian.Uint32(buffer)

		// read payload
		buffer = make([]byte, payloadSize)
		err = readBytes(reader, buffer)
		if err != nil {
			fmt.Printf(
				"read payload from connection %s failed: %s\n",
				conn.RemoteAddr().String(),
				err.Error(),
			)
			break
		}

		payload := string(buffer)
		fmt.Printf("received payload from connection %s: %s\n", conn.RemoteAddr().String(), payload)
	}
}

// read bytes from reader
func readBytes(reader io.Reader, buffer []byte) error {
	n, err := reader.Read(buffer)
	if err != nil {
		return err
	}

	if n != len(buffer) {
		return readBytes(reader, buffer[n:])
	}

	return nil
}
