package models
//用户表为untitled，约定大于配置原则
//tag设置结构体对用的数据表的列
type Untitled struct {
	//Id int64 `gorm:"column:id"`
	//Username string `gorm:"column:username"`
	//RealName string `gorm:"column:real_name"`
	//Age int `gorm:"column:age"`
	//Sex int `gorm:"column:sex"`
	//Address string `gorm:"column:address"`

	Id int64   `json:"id"`
	Username string  `json:"username"`
	RealName string  `json:"realname"`
	Age int   `json:"age"`
	Sex int   `json:"sex"`
	Address string  `json:"address"`
}