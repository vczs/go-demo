package my_server

import (
	"fmt"
	"net"
	"server/tool"
)

func serverWritePkg(conn net.Conn, mes []byte) error {

	//将int类型的mes的长度 len(mes) 转为[]byte类型
	mesLen := len(mes)
	mesLenBytes, mesLenBytesErr := tool.IntToBytes(mesLen)
	if mesLenBytesErr != nil {
		fmt.Printf("mesLenBytesErr：%v\n", mesLenBytesErr)
		return mesLenBytesErr
	}

	//先发送mes的长度给对方
	_, writeMesLenErr := conn.Write(mesLenBytes)
	if writeMesLenErr != nil {
		fmt.Printf("writeMesLenErr：%v\n", writeMesLenErr)
		return writeMesLenErr
	}

	//发送mes内容给对方
	_, writeMesErr := conn.Write(mes)
	if writeMesErr != nil {
		fmt.Printf("writeMesJsonErr：%v\n", writeMesErr)
		return writeMesErr
	}

	return nil
}
