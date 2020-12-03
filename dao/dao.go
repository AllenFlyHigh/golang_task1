package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"task1/models"
	snowFlake "task1/until"
)

var (
	db *gorm.DB
	err error
)

func init ()  {
	db, err = gorm.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	} else {
		// 全局禁用表名复数
		db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

		// 一般不会直接用CreateTable创建表
		// 检查模型`User`表是否存在，否则为模型`User`创建表
		if !db.HasTable(&models.Untitled{}) {
			if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Untitled{}).Error; err !=nil {
				panic(err)
			}
		}
	}
}
//查询所有的用户
func DBListUser(user *[] models.Untitled)  {
	db.Find(&user)
}
//新增用户
func DBCreateUser (untitled * models.Untitled) (tag bool) {
	var tempUser models.Untitled
	id := snowFlake.GetSnowflakeId() // 雪花算法生成唯一的id
	name := untitled.Username
	realName := untitled.RealName
	untitled.Id = id
	if name == "" || realName == "" {
		return false
	}
	db.First(&tempUser, id)
	if tempUser.Id == id  {
		fmt.Println("用户id已经被占用")
		return false
	} else {
		db.Create(&untitled)
		return true
	}
	return true
}
//条件查询
func DBQueryUsers (untitled *models.Untitled) (int64, [] models.Untitled) {
	username := untitled.Username
	sex := untitled.Sex
	age := untitled.Age

	var untitleds [] models.Untitled
	if username != "" {
		//根据用户名查询
		result := db.Where("username LIKE ?", "%" + username + "%").Find(&untitleds)
		return result.RowsAffected, untitleds
	} else {
		//年龄和性别的多条件查询
		if sex != 0 && age != 0 {
			result := db.Where("sex = ? AND age >= ?", sex, age).Find(&untitleds)
			return result.RowsAffected, untitleds
		}
	}
	return 0, untitleds
}
//删除用户
func DBDeleteUsers (id string) (models.Untitled)  {
	var untitled models.Untitled
	db.First(&untitled, id)
	if untitled.Id != 0 {
		db.Delete(&untitled)
		return untitled
	} else{
		return untitled
	}
}
//更新用户
func DBUpdateUsers (id string, bindUntitled *models.Untitled) (models.Untitled, error) {
	var untitled models.Untitled
	err := db.First(&untitled, id).Error
	fmt.Println(bindUntitled.Username)
	fmt.Println(bindUntitled.RealName)
	if err != nil {
		return untitled,err
	} else {
		bindUntitled.Id = untitled.Id
		db.Save(&bindUntitled)
		return *bindUntitled, err
	}
}
