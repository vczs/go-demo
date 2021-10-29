package message

const (
	LoginMesType      = "LoginMes"
	LoginResMesType   = "LoginResMes"
	RegisterMesType   = "RegisterMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"`//消息的类型
}

//登录
type LoginMes struct {
	Admin int `json:"admin"`
	Password string `json:"password"`
	AdminName string `json:"adminName"`
}

//登录结果
type LoginResMes struct {
	Code int  `json:"code"`//返回状态码
	Error string  `json:"error"`//返回错误信息
}

//注册
type RegisterMes struct {

}