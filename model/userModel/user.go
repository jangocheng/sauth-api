package userModel

import "time"

/*
	user 用户信息
*/
type User struct {
	Uuid       string    // 主键
	UserName   string    // 用户名
	Password   string    // 密码
	NickName   string    // 昵称
	Sex        string    // 性别
	Position   string    // 职位
	Tel        string    // 联系电话
	Mail       string    // 邮箱
	Address    string    // 地址
	Remark     string    // 备注
	CreateUser string    // 创建者
	CreateTime time.Time // 创建时间
	UpdateUser string    // 修改者
	UpdateTime time.Time // 修改时间
}
