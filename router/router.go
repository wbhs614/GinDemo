package router

import (
	"fmt"
	"ginDemo/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test() {
	fmt.Println("router了")
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.POST("/test", loginFunc)
		v1.POST("/test2", controllers.LoginTest)
		v1.POST("/getStudentList", controllers.GetStudentDetail)
		v1.POST("/addStudentInfo", controllers.AddStudentInfo)
		v1.POST("/updateStudentinfo", controllers.UpdateStudentInfo)
		v1.POST("/getStudentInfo", controllers.GetStudentInfoById)
		fmt.Println("进入路由啦")
	}

	r.Run(":9091")
}

func loginFunc(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	nameMap := make(map[string]interface{})
	if len(firstName) > 0 && len(lastName) > 0 {
		nameMap["firstname"] = firstName
		nameMap["lastname"] = lastName
		c.JSON(http.StatusOK, nameMap)
	} else {
		nameMap["message"] = "参数错误，请重新输入"
		nameMap["code"] = 2
		c.JSON(http.StatusOK, nameMap)
	}
}
