package Controller

import (
	"user-management-system/Model"
)

type UserController struct {
	userNumber int
	userArr    []Model.User
}

func NewUserController() *UserController {
	userController := &UserController{}
	userController.userNumber = 1
	newUser := Model.NewUser(userController.userNumber, "Jack", "男", 20, "10086", "杰克@163.com")
	userController.userArr = append(userController.userArr, newUser)
	return userController
}

func (userController *UserController) GetList() []Model.User {
	return userController.userArr
}

func (userController *UserController) Update(id int) *Model.User {
	index := userController.FindById(id)
	if index == -1 {
		return nil
	} else {
		return &userController.userArr[index]
	}
}

func (userController *UserController) Add(addUser Model.User) bool {
	userController.userNumber += 1
	addUser.Id = userController.userNumber
	userController.userArr = append(userController.userArr, addUser)
	return true
}

//根据用户Id查找用户在切片中的下标,没有就返回-1
func (userController *UserController) FindById(id int) int {
	index := -1
	for i, v := range userController.userArr {
		if v.Id == id {
			index = i
		}
	}
	return index
}

//根据Id在切片中删除用户
func (userController *UserController) Delete(id int) bool {
	index := userController.FindById(id)
	if index == -1 {
		return false
	} else {
		userController.userArr = append(userController.userArr[:index], userController.userArr[index+1:]...)
		return true
	}
}
