package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

// 这里我们定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 数据内容
}

// 定义两个消息.. 后面需要再增加

type LoginMes struct {
	UserId   int    `json:"userId"`   // 用户id
	UserPwd  string `json:"userPwd"`  // 用户名密码
	UserName string `json:"userName"` // 用户名
}

type LoginResMes struct {
	Code    int    `json:"code"` // 返回状态码 500 表示改用户未注册
	UsersId []int  // 增加字段，保存用户 id 的切片
	Error   string `json:"error"` // 返回错误信息
}

type RegisterMes struct {
	User User `json:"user"` // 类型就是User 结构体
}

type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回状态码 500 表示改用户未注册
	Error string `json:"error"` // 返回错误信息
}

// 为了配合服务器端推动用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` // 用户id
	Status int `json:"status"`
}

// 增加一个SmsMes 发送的消息
type SmsMes struct {
	Content string `json:"context"`
	User           // 匿名结构体， 继承
}
