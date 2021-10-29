package my_client

import (
	"client/tool"
	"errors"
	"fmt"
	"net"
)

func clientReadPkg(conn net.Conn) ([]byte, error) {

	//函数退出关闭链接
	defer func() {
		connCloseErr := conn.Close()
		if connCloseErr != nil {
			fmt.Printf("connCloseErr：%v\n", connCloseErr)
			return
		}
		fmt.Println("connClose sec...")
	}()

	buf := make([]byte, 1024) //创建一个切片接收读取的消息

	//首先接收对方发来信息的长度
	_, clientReadPkgLenErr := conn.Read(buf)
	if clientReadPkgLenErr != nil {
		fmt.Printf("clientReadPkgLenErr=%v\n", clientReadPkgLenErr)
		return nil, clientReadPkgLenErr
	}

	//读取到的内容存储在buf里 将buf转为int获取信息的长度
	mesReadLen, mesReadLenErr := tool.BytesToInt(buf)
	if mesReadLenErr != nil {
		fmt.Printf("mesReadLenErr=%v\n", mesReadLenErr)
		return nil, mesReadLenErr
	}

	//现在开始读取信息内容
	clientReadMes, clientReadMesErr := conn.Read(buf)
	if clientReadMesErr != nil {
		fmt.Printf("clientReadMesErr=%v\n", clientReadMesErr)
		return nil, clientReadMesErr
	}

	//判断接收的信息大小和刚才收到的大小是否相等
	if clientReadMes != mesReadLen {
		infoErr := fmt.Sprintf("信息包错误：{原信息包%v字节  当前信息包%v字节}", mesReadLen, clientReadMes)
		return nil, errors.New(infoErr)
	}

	return buf[:mesReadLen], nil
}
