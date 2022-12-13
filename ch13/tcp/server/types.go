package server

// request package
type RequestPackage struct {
	// request id
	Id string

	// command
	Command uint16

	// json parameters
	Parameters []byte
}

// response package
type ResponsePackage struct {
	// request id
	RequestId string

	// data
	Data any
}
