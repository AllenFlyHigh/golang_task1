package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task1/dao"
	"task1/models"
)

func Router ()  {
	// 链接 mysql
	router := gin.Default()
	// 路径映射
	router.GET("/stu/list", ListUser)
	router.POST("/stu/add", CreateUser)
	router.POST("/stu/query", GetUsers)
	router.PUT("/stu/update/:id", UpdateUser)
	router.DELETE("/stu/del/:id", DeleteUser)
	router.Run(":8080")
}
//列出所有的用户  127.0.0.1:8080/stu/list
func ListUser(c *gin.Context) {
	var untitled []models.Untitled
	dao.DBListUser(&untitled)
	c.JSON(http.StatusOK, &untitled)	//限制查找前line行
}
//创建用户
//curl -i -X POST -H "Content-Type: application/json" -d "{ \"username\": \"John\", \"realname\": \"Jack\"}" http: localhost:8080/stu/add
func CreateUser(c *gin.Context) {
	var untitled models.Untitled
	c.BindJSON(&untitled)	// 使用bindJson填充数据
	if dao.DBCreateUser(&untitled) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 200,
			"msg": "Fields are empty",
			"data":nil,
			})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"mag": "sucess",
			"data": &untitled,
		})
	}
}
//1.根据用户名模糊查询
//curl -i -X POST -H "Content-Type: application/json" -d "{\"username\":\"a\"}" http://localhost:8080/stu/query
// 2.年龄和性别的多条件查询
//curl -i -X POST -H "Content-Type: application/json" -d "{\"sex\":1,\"age\":20}" http://localhost:8080/stu/query
func GetUsers(c *gin.Context) {
	var untitled models.Untitled   //传入的参数
	c.BindJSON(&untitled)	// 使用bindJson填充数据
	count, untitleds := dao.DBQueryUsers(&untitled)
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code" : 200,
			"msg" : "sucess",
			"data": &untitleds,
			"count": count,//count用来分页
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg" : "success",
			"data" : nil,
		})
	}
}
//删除用户
//curl -i -X DELETE http://localhost:8080/stu/del/1
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	if untitled := dao.DBDeleteUsers(id); untitled.Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": &untitled,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "success",
			"data": nil,
		})
	}
}
//更新用户

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var untitled models.Untitled
	c.BindJSON(&untitled)
	untitled, err := dao.DBUpdateUsers(id, &untitled)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg": "not sucess",
			"data": nil,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 200,
			"msg": "sucess",
			"data": &untitled,
		})
	}
}

