package model

type User struct {
	Id    int
	Name  string
	Age   int
	State int
}

// 配置操作数据库的表名称
func (user User) TableName() string {
	return "user"
}
