package my_client

import (
	"client/message"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	loop   bool
	option int
)

func StartExe() {
	showHomeMenu()
}

func showHomeMenu() {
	loop = true
	for loop {
		fmt.Println("***用户通讯系统***")
		fmt.Println("*     1.登录     *")
		fmt.Println("*     2.注册     *")
		fmt.Println("*     3.退出     *")
		fmt.Println("*****************")
		selectOption()
	}
}

func selectOption() {
	fmt.Println("请选择(1-3):")
	_, selectOptionErr := fmt.Scanln(&option)
	if selectOptionErr != nil {
		fmt.Printf("selectOptionErr:%v\n", selectOptionErr)
		return
	}
	switch option {
	case 1:
		signInOption()
		loop = false
	case 2:
		fmt.Println("2")
		loop = false
	case 3:
		fmt.Println("3")
		loop = false
	default:
		fmt.Println("输入有误，请重新输入。")
	}
}

func signInOption() {

	var admin int
	var password string

	fmt.Print("请输入帐号：")
	_, adminInputErr := fmt.Scanln(&admin)
	if adminInputErr != nil {
		fmt.Printf("selectOptionErr:%v\n", adminInputErr)
		return
	}

	fmt.Print("请输入密码：")
	_, passwordInputErr := fmt.Scanln(&password)
	if passwordInputErr != nil {
		fmt.Printf("selectOptionErr:%v\n", passwordInputErr)
		return
	}

	//将用户输入的信息 生成json
	clientConn, signInErr := signIn(admin, password)
	if signInErr != nil {
		fmt.Printf("signInErr:%v\n", signInErr)
		return
	}

	clientRead, clientReadPkgErr := clientReadPkg(clientConn)
	if clientReadPkgErr != nil {
		fmt.Printf("clientReadPkgErr:%v\n", clientReadPkgErr)
		return
	}

	loginErr := loginResMesProcess(clientRead)
	if loginErr != nil {
		fmt.Printf("loginErr:%v\n", loginErr)
		return
	}

	fmt.Println("登陆成功！！！")
}

//处理登陆结果信息包判断是否登录成功
func loginResMesProcess(clientRead []byte) error {

	var mes message.Message
	mesUnmarshalErr := json.Unmarshal(clientRead, &mes)
	if mesUnmarshalErr != nil {
		fmt.Printf("mesUnmarshalErr:%v\n", mesUnmarshalErr)
		return mesUnmarshalErr
	}

	var loginResMes message.LoginResMes
	loginResUnmarshalErr := json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResUnmarshalErr != nil {
		fmt.Printf("loginResUnmarshalErr:%v\n", loginResUnmarshalErr)
		return loginResUnmarshalErr
	}

	if loginResMes.Code == 200 {
		return nil
	} else if loginResMes.Code == 500 {
		loginErr := errors.New(loginResMes.Error)
		return loginErr
	} else {
		loginErr := errors.New("未知错误！！！")
		return loginErr
	}
}
