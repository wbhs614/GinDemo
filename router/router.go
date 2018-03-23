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
		//学生信息的接口
		v1.POST("/student/getStudentList", controllers.GetStudentDetail)
		v1.POST("/student/addStudentInfo", controllers.AddStudentInfo)
		v1.POST("/student/updateStudentinfo", controllers.UpdateStudentInfo)
		v1.POST("/student/getStudentInfo", controllers.GetStudentInfoById)
		//课程的接口
		v1.POST("/course/addCourse", controllers.AddCourse)
		v1.POST("/course/getCourseList", controllers.GetCourseList)
		v1.POST("/course/addCoursePlan", controllers.AddCoursePlane)
		v1.POST("/course/getCoursePlanList", controllers.GetCoursePlanList)
		v1.POST("/course/addSelectCoursePlan", controllers.AddSelectCoursePlan)
		v1.POST("/course/getStListByChooseid", controllers.GetInfoListByChooseid)
		v1.POST("/course/getChListByStudentid", controllers.GetInfoListByStudentid)

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
