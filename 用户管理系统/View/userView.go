package View

import (
	"fmt"
	"strconv"
	"user-management-system/Controller"
	"user-management-system/Model"
)

type userView struct {
	userController *Controller.UserController
	loop           bool
	key            string
}

func Start() {
	userView := &userView{
		loop: true,
	}
	userView.userController = Controller.NewUserController()
	userView.showMenu()
}

func (userView *userView) showMenu() {
	for userView.loop {
		fmt.Println("***用户管理系统***")
		fmt.Println("*     1.添加     *")
		fmt.Println("*     2.删除     *")
		fmt.Println("*     3.查看     *")
		fmt.Println("*     4.修改     *")
		fmt.Println("*     5.退出     *")
		fmt.Println("*****************")
		userView.Select()
	}
}

//选择功能
func (userView *userView) Select() {
	fmt.Println("请选择(1-5):")
	fmt.Scanln(&userView.key)
	switch userView.key {
	case "1":
		userView.add()
	case "2":
		userView.delete()
	case "3":
		userView.list()
	case "4":
		userView.update()
	case "5":
		userView.signOut()
	default:
		fmt.Println("输入有误！！！")
	}
}

//修改
func (userView *userView) update() {
	fmt.Println("-----------------------修改用户-----------------------")
	fmt.Println("请输入要修改的用户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("已退出修改用户功能。。。")
		return
	}
	user := userView.userController.Update(id)
	if user != nil {
		fmt.Println("请输入要修改的内容，直接回车表示不修改。")
		fmt.Printf("姓名(%v):", user.Name)
		var name string
		fmt.Scanln(&name)
		if name != "" {
			user.Name = name
		}
		fmt.Printf("性别(%v):", user.Gender)
		var gender string
		fmt.Scanln(&gender)
		if name != "" {
			user.Gender = gender
		}
		fmt.Printf("年龄(%v):", user.Age)
		var age string
		fmt.Scanln(&age)
		if age != "" {
			data, _ := strconv.Atoi(age)
			user.Age = int(data)
		}
		fmt.Printf("电话(%v):", user.Phone)
		var phone string
		fmt.Scanln(&phone)
		if phone != "" {
			user.Phone = phone
		}
		fmt.Printf("邮箱(%v):", user.Email)
		var email string
		fmt.Scanln(&email)
		if email != "" {
			user.Email = email
		}
		fmt.Println("-----------------------修改成功-----------------------")
	} else {
		fmt.Printf("编号\"%v\"用户不存在！！！\n", id)
		fmt.Println("-----------------------修改失败-----------------------")
	}
}

//删除
func (userView *userView) delete() {
	fmt.Println("-----------------------删除用户-----------------------")
	fmt.Println("请输入要删除的用户编号(-1退出):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		fmt.Println("已退出删除用户功能。。。")
		return
	}
	ok := ""
	fmt.Println("确认删除(输入\"Y\"确认删除,其他则取消删除):")
	fmt.Scanln(&ok)
	if ok == "y" || ok == "Y" {
		if userView.userController.Delete(id) {
			fmt.Printf("已删除编号\"%v\"用户！！！\n", id)
			fmt.Println("-----------------------删除成功-----------------------")
		} else {
			fmt.Printf("编号\"%v\"用户不存在！！！\n", id)
			fmt.Println("-----------------------删除失败-----------------------")
		}
	} else {
		fmt.Printf("已取消删除！！！\n")
		fmt.Println("-----------------------取消删除-----------------------")
	}
}

//查看
func (userView *userView) list() {
	userArr := userView.userController.GetList()
	fmt.Println("-----------------------用户列表-----------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱\t")
	for i := 0; i < len(userArr); i++ {
		fmt.Println(userArr[i].GetInfo())
	}
	fmt.Println("-------------------------END-------------------------")
}

//增加
func (userView *userView) add() {
	fmt.Println("-----------------------添加用户-----------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	newUser := Model.NewUser2(name, gender, age, phone, email)
	if userView.userController.Add(newUser) {
		fmt.Println("-----------------------添加成功-----------------------")
	} else {
		fmt.Println("-----------------------添加失败-----------------------")
	}
}

//退出程序
func (userView *userView) signOut() {
	ok := ""
	fmt.Println("是否退出程序(y/n):")
	for {
		fmt.Scanln(&ok)
		if ok == "n" || ok == "y" || ok == "N" || ok == "Y" {
			break
		}
		fmt.Println("输入有误,请重新输入(y/n):")
	}
	if ok == "y" || ok == "Y" {
		userView.loop = false
		fmt.Println("程序已退出。。。")
	}
}
