package my_client

import (
	"fmt"
	"net"
)

func connectServer() (net.Conn, error) {
	conn, dialErr := net.Dial("tcp", "127.0.0.1:7000")
	if dialErr != nil {
		fmt.Printf("dialErr：%v\n", dialErr)
		return nil, dialErr
	}
	return conn, nil
}
