package my_server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"server/message"
	"server/tool"
)

func StartExe() {

	//开启监听
	fmt.Println("The 服务器 starts listening on port 7000 ...")
	listen, listenErr := net.Listen("tcp", "0.0.0.0:7000")
	if listenErr != nil {
		fmt.Printf("listenErr：%v\n", listenErr)
		return
	}

	//函数退出后关闭listen
	defer func() {
		listenCloseErr := listen.Close() //关闭listen
		if listenCloseErr != nil {
			fmt.Printf("listenCloseErr=%v\n", listenCloseErr)
			return
		}
		fmt.Println("listenClose suc...")
	}()

	//监听开启成功 等待客户端连接
	for {
		conn, acceptErr := listen.Accept()
		if acceptErr != nil {
			fmt.Printf("acceptErr=%v\n", acceptErr)
		} else {
			ip := conn.RemoteAddr().String()
			fmt.Printf("New: {%v} connect ...\n", ip)
			go Process(conn, ip)
		}
	}
}

func Process(conn net.Conn, ip string) {

	defer func() {
		connCloseErr := conn.Close() //关闭listen
		if connCloseErr != nil {
			fmt.Printf("connCloseErr=%v\n", connCloseErr)
			return
		}
		fmt.Println("connClose suc...")
	}()

	for {
		serverReadPkg, serverReadPkgErr := serverReadPkg(conn, ip)
		if serverReadPkgErr != nil {
			fmt.Printf("serverReadPkgErr=%v\n", serverReadPkgErr)
			return
		}
		fmt.Printf("json:%v\n", string(serverReadPkg))

		//反序列化serverReadPkg
		var mes message.Message
		mesUnmarshalErr := json.Unmarshal(serverReadPkg, &mes)
		if mesUnmarshalErr != nil {
			fmt.Printf("mesUnmarshalErr:%v", mesUnmarshalErr)
			return
		}

		//序列化完成传给mesProcess进行处理
		mesProcessErr := mesProcess(conn, &mes)
		if mesProcessErr != nil {
			fmt.Printf("mesProcessErr:%v", mesProcessErr)
			return
		}
	}

}

func serverReadPkg(conn net.Conn, ip string) ([]byte, error) {

	buf := make([]byte, 1024) //创建一个切片接收读取的消息

	//首先接收客户端发来信息的长度
	_, serverReadMesLenErr := conn.Read(buf)
	if serverReadMesLenErr != nil {
		if serverReadMesLenErr == io.EOF {
			fmt.Printf("{%v}断开连接！\n", ip)
		} else {
			fmt.Printf("serverReadMesLenErr:%v\n", serverReadMesLenErr)
		}
		return nil, serverReadMesLenErr
	}

	//读取到的内容存储在buf里 将buf转为int获取信息的长度
	mesReadLen, mesReadLenErr := tool.BytesToInt(buf)
	if mesReadLenErr != nil {
		fmt.Printf("mesReadLenErr=%v\n", mesReadLenErr)
		return nil, mesReadLenErr
	}

	//现在开始读取信息内容
	serverReadMes, serverReadMesErr := conn.Read(buf)
	if serverReadMesErr != nil {
		if serverReadMesErr == io.EOF {
			fmt.Printf("{%v}断开连接！\n", ip)
		} else {
			fmt.Printf("serverReadMesErr:%v\n", serverReadMesErr)
		}
		return nil, serverReadMesErr
	}
	//判断接收的信息大小和刚才收到的大小是否相等
	if serverReadMes != mesReadLen {
		infoErr := fmt.Sprintf("信息包错误：{原信息包%v字节  当前信息包%v字节}", mesReadLen, serverReadMes)
		return nil, errors.New(infoErr)
	}
	return buf[:mesReadLen], nil
}
