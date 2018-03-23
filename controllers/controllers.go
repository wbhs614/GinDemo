package controllers

import (
	"fmt"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func LoginTest(c *gin.Context) {
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

func AddStudentInfo(c *gin.Context) {
	dataMap := make(map[string]interface{})
	name := c.Request.FormValue("name")
	age := c.PostForm("age")
	sex := c.PostForm("sex")
	address := c.PostForm("address")
	phone := c.PostForm("phone")
	guardian := c.PostForm("guardian")
	grade := c.PostForm("grade")
	class := c.PostForm("class")
	if len(name) == 0 || len(age) == 0 || len(sex) == 0 || len(address) == 0 || len(phone) == 0 || len(guardian) == 0 {
		dataMap["code"] = 2
		dataMap["message"] = "请填写必填项"
	} else {
		db, err := utils.OpenGinDB()
		if err != nil {
			dataMap["code"] = 4
			dataMap["message"] = "查询数据库失败"
		} else {
			lowSex := strings.ToLower(sex)
			if lowSex != "f" && lowSex != "m" {
				dataMap["code"] = 6
				dataMap["message"] = "性别格式不对"
				c.JSON(http.StatusOK, dataMap)
				return
			}
			params := make(map[string]interface{})
			params["name"] = name
			params["age"] = age
			params["sex"] = sex
			params["address"] = address
			params["phone"] = phone
			params["guardian"] = guardian
			params["grade"] = grade
			params["class"] = class
			params["studentid"] = utils.CreateStudentid()
			_, err := db.Table("studentInfo").Data(params).Insert()
			if err != nil {
				dataMap["code"] = 3
				dataMap["errInfo"] = err
				dataMap["message"] = "添加数据失败"
			} else {
				dataMap["code"] = 0
				dataMap["message"] = "添加数据成功"
			}
		}
		db.Close()
	}
	c.JSON(http.StatusOK, dataMap)
}

func GetStudentDetail(c *gin.Context) {
	dataMap := make(map[string]interface{})
	limitstr := c.PostForm("limit")
	offsetstr := c.PostForm("offset")
	if len(limitstr) == 0 || len(offsetstr) == 0 {
		limitstr = "20"
		offsetstr = "0"
	}
	db, err := utils.OpenGinDB()
	if err != nil {
		dataMap["code"] = 4
		dataMap["message"] = "查询数据库失败"
	} else {
		dataMap := make(map[string]interface{})
		list, err := db.Query("select * from  studentInfo LIMIT ? OFFSET ?", limitstr, offsetstr)
		if err != nil {
			dataMap["code"] = 4
			dataMap["message"] = "查询数据库失败"
		} else {
			count := len(list)
			if count > 0 {
				dataMap["data"] = list
				dataMap["code"] = 0
				dataMap["message"] = "获取数据成功"
			} else {
				dataMap["code"] = 2
				dataMap["message"] = "暂无数据"
			}
		}
		c.JSON(http.StatusOK, dataMap)
	}
	db.Close()
	c.JSON(http.StatusOK, dataMap)
}

func UpdateStudentInfo(c *gin.Context) {
	dataMap := make(map[string]interface{})
	studentId := c.PostForm("studentid")
	if len(studentId) == 0 {
		dataMap["code"] = 4
		dataMap["message"] = "请输入必填项"
	} else {
		db, err := utils.OpenGinDB()
		where := fmt.Sprintf("studentid=%s", studentId)
		studentTable := db.Table("studentInfo")
		count, err := studentTable.Where(where).Count("*")
		if err != nil {
			dataMap["code"] = 3
			dataMap["message"] = "查询数据库失败"
		} else {
			if count > 0 {
				params := make(map[string]interface{})
				name := c.PostForm("name")
				age := c.PostForm("age")
				sex := c.PostForm("sex")
				address := c.PostForm("address")
				phone := c.PostForm("phone")
				guardian := c.PostForm("guardian")
				grade := c.PostForm("grade")
				class := c.PostForm("class")
				if len(name) > 0 {
					params["name"] = name
				}
				if len(age) > 0 {
					intage, _ := strconv.ParseInt(age, 10, 64)
					params["age"] = intage
				}
				if len(sex) > 0 {
					lowSex := strings.ToLower(sex)
					if lowSex != "f" && lowSex != "m" {
						dataMap["code"] = 6
						dataMap["message"] = "性别格式不对"
						c.JSON(http.StatusOK, dataMap)
						return
					}
					params["sex"] = sex
				}
				if len(address) > 0 {
					params["address"] = address
				}
				if len(phone) > 0 {
					params["phone"] = phone
				}
				if len(guardian) > 0 {
					params["guardian"] = guardian
				}
				if len(grade) > 0 {
					intgrade, _ := strconv.ParseInt(grade, 10, 64)
					params["grade"] = intgrade
				}
				if len(class) > 0 {
					intclass, _ := strconv.ParseInt(class, 10, 64)
					params["class"] = intclass
				}
				fmt.Println(params)
				if len(params) > 0 {
					fmt.Println("hehhe")
					_, err := studentTable.Where("studentid", studentId).Data(params).Update()
					if err == nil {
						dataMap["code"] = 0
						dataMap["message"] = "更新学生信息成功"
					} else {
						dataMap["code"] = 3
						dataMap["message"] = "更新学生信息失败"
					}
				} else {
					dataMap["code"] = 5
					dataMap["message"] = "请输入需要修改的信息"
				}

			} else {
				dataMap["code"] = 4
				dataMap["message"] = "该学生不存在"
			}
		}
	}
	c.JSON(http.StatusOK, dataMap)
}

func GetStudentInfoById(c *gin.Context) {
	studentId := c.PostForm("studentid")
	dataMap := make(map[string]interface{})
	if len(studentId) == 0 {
		dataMap["code"] = 2
		dataMap["message"] = "参数错误：缺少studentId"
	} else {
		db, err := utils.OpenGinDB()
		where := fmt.Sprintf("studentid=%s", studentId)
		studentTable := db.Table("studentInfo")
		st, err := studentTable.Where(where).First()
		if err != nil {
			dataMap["code"] = 3
			dataMap["message"] = "查询学生信息失败"
		} else {
			if len(st) == 0 {
				dataMap["code"] = 4
				dataMap["message"] = "没有查询到学生信息"
			} else {
				dataMap["code"] = 0
				dataMap["message"] = "获取学生信息成功"
				dataMap["data"] = st
			}
		}
		db.Close()
	}
	c.JSON(http.StatusOK, dataMap)
}
