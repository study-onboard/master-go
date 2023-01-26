package server

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

// package struct
// -------------------------------------------
// uint16 - command
// uint32 - parameters size
// bytes[size] - parameters json

// read package
func readRequestPackage(reader io.Reader, conn net.Conn) (request *RequestPackage, err error) {
	// read id -- uuid
	buffer := make([]byte, 36)
	err = readBytes(reader, buffer)
	if err != nil {
		fmt.Printf(
			"read request id from connection %s failed: %s\n",
			conn.RemoteAddr().String(),
			err.Error(),
		)
		return
	}
	id := string(buffer)
	fmt.Printf(
		"received package request id from connection %s: %s\n",
		conn.RemoteAddr().String(),
		id,
	)

	// read command
	buffer = make([]byte, 2)
	err = readBytes(reader, buffer)
	if err != nil {
		fmt.Printf(
			"read request command from connection %s failed: %s\n",
			conn.RemoteAddr().String(),
			err.Error(),
		)
	}
	command := binary.BigEndian.Uint16(buffer)
	fmt.Printf(
		"received package command from connection %s: %d\n",
		conn.RemoteAddr().String(),
		command,
	)

	// read parameters size
	buffer = make([]byte, 4)
	err = readBytes(reader, buffer)
	if err != nil {
		fmt.Printf(
			"read payload size from connection %s failed: %s\n",
			conn.RemoteAddr().String(),
			err.Error(),
		)
		return
	}
	parametersSize := binary.BigEndian.Uint32(buffer)

	// read parameters
	buffer = make([]byte, parametersSize)
	err = readBytes(reader, buffer)
	if err != nil {
		fmt.Printf(
			"read payload from connection %s failed: %s\n",
			conn.RemoteAddr().String(),
			err.Error(),
		)
		return
	}

	// return result
	err = nil
	request = &RequestPackage{
		Id:         id,
		Command:    command,
		Parameters: buffer,
	}

	return
}

// write response package
func writeResponsePackage(response *ResponsePackage, writer io.Writer, conn net.Conn) error {
	// write request id
	err := writeBytes(writer, []byte(response.RequestId))
	if err != nil {
		fmt.Printf(
			"write response request id to %s failed: %s\n",
			conn.RemoteAddr().String(),
			err.Error(),
		)
		return err
	}

	// encode data
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.Encode(response.Data)
	data := buffer.Bytes()

	// write data size
	dataSize := uint32(len(data))
	buffer = bytes.Buffer{}
	binary.Write(&buffer, binary.BigEndian, dataSize)
	err = writeBytes(writer, buffer.Bytes())
	if err != nil {
		fmt.Printf(
			"write response data size to %s failed: %s\n",
			conn.RemoteAddr().String(),
			err.Error(),
		)
		return err
	}

	// write data
	err = writeBytes(writer, data)
	if err != nil {
		fmt.Printf(
			"write response data to %s failed: %s\n",
			conn.RemoteAddr().String(),
			err.Error(),
		)
		return err
	}

	return nil
}

// write bytes by writer
func writeBytes(writer io.Writer, buffer []byte) error {
	n, err := writer.Write(buffer)
	if err != nil {
		return err
	}

	if n != len(buffer) {
		return writeBytes(writer, buffer[n:])
	}

	return nil
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
