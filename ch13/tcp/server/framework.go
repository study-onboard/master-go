package server

import "net"

// handle request
func handleRequest(request *RequestPackage, conn net.Conn) (*ResponsePackage, error) {
	return &ResponsePackage{
		RequestId: request.Id,
		Data: map[string]any{
			"name":  "Kut Zhang",
			"level": 33,
		},
	}, nil
}
