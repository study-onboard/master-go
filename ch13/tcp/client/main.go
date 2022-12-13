package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/google/uuid"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// write request id
	id := uuid.NewString()
	_, err = writer.Write([]byte(id))
	if err != nil {
		panic(err)
	}

	// write command
	command := uint16(1)
	commandBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(commandBytes, command)
	_, err = writer.Write(commandBytes)
	if err != nil {
		panic(err)
	}

	// write parameters
	parameters := map[string]any{
		"name": "Kut Zhang",
		"age":  99,
	}

	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	encoder.Encode(parameters)
	data := b.Bytes()

	// write parameters size
	dataSizeBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(dataSizeBytes, uint32(len(data)))
	_, err = writer.Write(dataSizeBytes)
	if err != nil {
		panic(err)
	}

	// write parameters data
	_, err = writer.Write(data)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	/////////////////////////////////////////////

	idBytes := make([]byte, 36)
	_, err = reader.Read(idBytes)
	if err != nil {
		panic(err)
	}
	id = string(idBytes)

	dataSizeBytes = make([]byte, 4)
	_, err = reader.Read(dataSizeBytes)
	if err != nil {
		panic(err)
	}

	dataSize := binary.BigEndian.Uint32(dataSizeBytes)
	data = make([]byte, dataSize)
	_, err = reader.Read(data)
	if err != nil {
		panic(err)
	}
	text := string(data)
	fmt.Printf("response: %s\n", text)

	conn.Close()
}
