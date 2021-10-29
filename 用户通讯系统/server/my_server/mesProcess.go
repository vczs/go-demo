package my_server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"server/message"
)

func mesProcess(conn net.Conn, mes *message.Message) (mesProcessErr error) {
	switch mes.Type {
	case message.LoginMesType:
		loginProcessErr := loginProcess(conn, mes)
		if loginProcessErr != nil {
			fmt.Printf("loginProcessErr:%v", loginProcessErr)
			return loginProcessErr
		}
	case message.RegisterMesType:
		//处理注册
	default:
		mesProcessErr = errors.New(fmt.Sprintln("无法识别消息类型"))
	}
	return mesProcessErr
}

func loginProcess(conn net.Conn, mes *message.Message) error {
	//处理登录
	loginMesProcessPkg, loginMesProcessPkgErr := loginMesProcess(mes)
	if loginMesProcessPkgErr != nil {
		fmt.Printf("loginMesProcessPkgErr:%v", loginMesProcessPkgErr)
		return loginMesProcessPkgErr
	}
	//发送结果
	serverWritePkgErr := serverWritePkg(conn, loginMesProcessPkg)
	if serverWritePkgErr != nil {
		fmt.Printf("serverWritePkgErr:%v", serverWritePkgErr)
		return serverWritePkgErr
	}
	return nil
}

//处理登录
func loginMesProcess(mes *message.Message) (loginMesProcessPkg []byte, loginMesProcessErr error) {

	//取出message.Data反序列化为LoginMes
	var loginMes message.LoginMes
	loginMesUnmarshalErr := json.Unmarshal([]byte(mes.Data), &loginMes)
	if loginMesUnmarshalErr != nil {
		fmt.Printf("loginMesUnmarshalErr:%v", loginMesUnmarshalErr)
		return nil, loginMesUnmarshalErr
	}

	//创建登录结果信息包详细信息
	var logResMes message.LoginResMes
	//判断LoginMes里的Admin和Password字段
	if loginMes.Admin == 123 && loginMes.Password == "abc" {
		//如果账户和密码正确  就给logResMes的字段赋值登陆成功的内容
		logResMes.Code = 200 //状态码200 错误信息为空表示登陆成功
	} else {
		//如果账户和密码不正确  就给logResMes的字段赋值登陆失败的内容
		logResMes.Code = 500     //状态码500 表示登陆失败
		logResMes.Error = "密码错误" //登陆失败 错误信息不能为空
	}
	//序列化logResMes
	logResMesJson, logResMesJsonErr := json.Marshal(logResMes)
	if logResMesJsonErr != nil {
		fmt.Printf("logResMesJsonErr:%v", logResMesJsonErr)
		return nil, logResMesJsonErr
	}

	//创建登陆结果信息包
	var mesRes message.Message
	mesRes.Type = message.LoginResMesType
	mesRes.Data = string(logResMesJson)

	//系列化mesRes
	mesResJson, mesResJsonErr := json.Marshal(mesRes)
	if mesResJsonErr != nil {
		fmt.Printf("mesResJsonErr:%v", mesResJsonErr)
		return nil, mesResJsonErr
	}

	return mesResJson, nil
}
