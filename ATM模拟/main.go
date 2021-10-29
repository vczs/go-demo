package main

import "fmt"

func main() {
	start(true, 10000).ShowMenu()
}

type MyAccount struct {
	loop    bool    //程序运行状态
	key     string  //用户输入的选项
	balance float64 //账户余额
	details string  //交易记录
	money   float64 //交易金额
	note    string  //交易说明
}

//构造函数
func start(loop bool, balance float64) *MyAccount {
	return &MyAccount{
		loop:    loop,
		balance: balance,
	}
}

//显示主菜单
func (myAccount *MyAccount) ShowMenu() {
	for myAccount.loop {
		fmt.Println("******欢迎使用ATM******")
		fmt.Println("*      1.交易记录      *")
		fmt.Println("*      2.转入金额      *")
		fmt.Println("*      3.转出金额      *")
		fmt.Println("*      4.退出程序      *")
		fmt.Println("***********************")
		myAccount.Select()
	}
}

//选择功能
func (myAccount *MyAccount) Select() {
	fmt.Println("请选择(1-4):")
	fmt.Scanln(&myAccount.key)
	switch myAccount.key {
	case "1":
		myAccount.Record()
	case "2":
		myAccount.Income()
	case "3":
		myAccount.Payment()
	case "4":
		myAccount.SignOut()
	default:
		fmt.Println("输入有误！！！")
	}
}

//交易明细
func (myAccount *MyAccount) Record() {
	if myAccount.details == "" {
		fmt.Println("无交易记录~~~")
	} else {
		fmt.Print("类型\t金额\t余额\t备注")
		fmt.Println(myAccount.details)
	}
}

//转入
func (myAccount *MyAccount) Income() {
	fmt.Println("转入金额:")
	fmt.Scanln(&myAccount.money)
	fmt.Println("备注:")
	fmt.Scanln(&myAccount.note)
	myAccount.balance += myAccount.money
	myAccount.details += fmt.Sprintf("\n转入\t+%v\t%v\t%v", myAccount.money, myAccount.balance, myAccount.note)
}

//转出
func (myAccount *MyAccount) Payment() {
	fmt.Println("转出金额:")
	fmt.Scanln(&myAccount.money)
	if myAccount.money > myAccount.balance {
		fmt.Println("余额不足！！！")
	} else {
		fmt.Println("备注:")
		fmt.Scanln(&myAccount.note)
		myAccount.balance -= myAccount.money
		myAccount.details += fmt.Sprintf("\n转出\t-%v\t%v\t%v", myAccount.money, myAccount.balance, myAccount.note)
	}
}

//退出程序
func (myAccount *MyAccount) SignOut() {
	ok := ""
	fmt.Println("退出程序(y/n):")
	for {
		fmt.Scanln(&ok)
		if ok == "n" || ok == "y" || ok == "N" || ok == "Y" {
			break
		} else {
			fmt.Println("输入有误,请重新输入(y/n):")
		}
	}
	if ok == "y" || ok == "Y" {
		myAccount.loop = false
		fmt.Println("程序已退出。。。")
	}
}
