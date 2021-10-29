package Model

import "fmt"

type User struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//User的构造函数
func NewUser(id int, name string, gender string, age int, phone string, email string) User {
	return User{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//User的构造函数
func NewUser2(name string, gender string, age int, phone string, email string) User {
	return User{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (user User) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", user.Id, user.Name, user.Gender, user.Age, user.Phone, user.Email)
	return info
}
