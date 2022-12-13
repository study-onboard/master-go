package server

import (
	"bufio"
	"fmt"
	"net"
)

// run server
func Run(port int) {
	fmt.Printf("Starting server on port %d........\n", port)

	l, err := net.Listen("tcp4", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server on port %d started!\n", port)

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
func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("handle connection from %s\n", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		// read request package
		request, err := readRequestPackage(reader, conn)
		if err != nil {
			fmt.Printf(
				"read request package from %s failed: %s\n",
				conn.RemoteAddr().String(),
				err.Error(),
			)
			break
		}

		// this err is business error,
		// when backend need break this connection,
		// reutrn err != nil
		response, err := handleRequest(request, conn)
		if err != nil {
			fmt.Printf(
				"handle request error from %s - %s: %s",
				conn.RemoteAddr().String(),
				request.Id,
				err.Error(),
			)
			break
		}

		// write response package
		err = writeResponsePackage(response, writer, conn)
		if err != nil {
			fmt.Printf(
				"write response package to %s failed: %s\n",
				conn.RemoteAddr().String(),
				err.Error(),
			)
			break
		}
		err = writer.Flush()
		if err != nil {
			fmt.Printf(
				"write flush response package to %s failed: %s\n",
				conn.RemoteAddr().String(),
				err.Error(),
			)
			break
		}
	}
}
