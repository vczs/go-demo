package my_client

import (
	"client/message"
	"client/tool"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func signIn(admin int, password string) (net.Conn, error) {

	//连接到服务器
	conn, connServerErr := connectServer()
	if connServerErr != nil {
		fmt.Printf("connServerErr：%v\n", connServerErr)
		return nil, connServerErr
	}
	//拿到conn 发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	var logMes message.LoginMes
	logMes.Admin = admin
	logMes.Password = password

	//将logMes序列化后赋值给mes的Data字段
	logMesJson, logMesJsonErr := json.Marshal(logMes)
	if logMesJsonErr != nil {
		fmt.Printf("logMesJsonErr：%v\n", logMesJsonErr)
		return nil, logMesJsonErr
	}
	mes.Data = string(logMesJson)

	//将mes序列化
	mesJson, mesJsonErr := json.Marshal(mes)
	if mesJsonErr != nil {
		fmt.Printf("mesJsonErr：%v\n", mesJsonErr)
		return nil, mesJsonErr
	}

	//用len(mesJson)获取mesJson的 int类型的长度
	mesJsonLen := len(mesJson)
	//将int类型的mesJsonLen转为[]byte类型
	mesJsonLenEnd, mesJsonLenEndErr := tool.IntToBytes(mesJsonLen)
	if mesJsonLenEndErr != nil {
		fmt.Printf("mesJsonLenEndErr：%v\n", mesJsonLenEndErr)
		return nil, mesJsonLenEndErr
	}

	//先发送mesJson的长度给服务端
	_, writeMesJsonLenErr := conn.Write(mesJsonLenEnd)
	if writeMesJsonLenErr != nil {
		fmt.Printf("writeMesJsonLenErr：%v\n", writeMesJsonLenErr)
		return nil, writeMesJsonLenErr
	}

	time.Sleep(time.Duration(1) * time.Microsecond) //休息一会  防止消息长度和内容连续发送无间隔 导致服务端无法区别

	//发送mesJson给服务器
	_, writeMesJsonErr := conn.Write(mesJson)
	if writeMesJsonErr != nil {
		fmt.Printf("writeMesJsonErr：%v\n", writeMesJsonErr)
		return nil, writeMesJsonErr
	}

	return conn, nil
}
