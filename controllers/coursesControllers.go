package controllers

import (
	"fmt"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	// "strings"
)

func AddCourse(c *gin.Context) {
	courseName := c.PostForm("coursename")
	isrequire := c.PostForm("isRequired")
	grade := c.PostForm("grade")
	profession := c.PostForm("profession")
	dataMap := make(map[string]interface{})
	if len(courseName) == 0 || len(isrequire) == 0 || len(grade) == 0 || len(profession) == 0 {
		dataMap["code"] = 2
		dataMap["message"] = "请填写完整必填项"
	} else {
		params := make(map[string]interface{})
		params["courseName"] = courseName
		intisrequire, _ := strconv.ParseInt(isrequire, 10, 64)
		params["isrequired"] = intisrequire
		intgrade, gerr := strconv.ParseInt(grade, 10, 64)
		if gerr != nil {
			intgrade = 0
		}
		params["grade"] = intgrade
		params["profession"] = profession
		params["courseid"] = "JSJ" + utils.CreateStudentid()
		db, err := utils.OpenGinDB()
		if err != nil {
			dataMap["code"] = 4
			dataMap["errInfo"] = err
			dataMap["message"] = "插入数据库失败"
		} else {
			_, inserterr := db.Table("course").Data(params).Insert()
			if inserterr != nil {
				dataMap["code"] = 5
				dataMap["errInfo"] = inserterr
				dataMap["message"] = "插入数据库失败"
			} else {
				dataMap["code"] = 0
				dataMap["message"] = "添加课程成功"
			}
		}
	}
	c.JSON(http.StatusOK, dataMap)
}

func GetCourseList(c *gin.Context) {
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
		dataMap["message"] = "打开数据库失败"
	} else {
		intlimit, limiterr := strconv.Atoi(limitstr)
		intoffset, offseterr := strconv.Atoi(offsetstr)
		if limiterr != nil {
			intlimit = 20
		}
		if offseterr != nil {
			intoffset = 20
		}
		list, cerr := db.Table("course").Fields("courseName,courseid,grade,isrequired,profession").Limit(intlimit).Offset(intoffset).Get()
		if cerr != nil {
			dataMap["code"] = 2
			dataMap["errinfo"] = cerr
			dataMap["message"] = "查询数据库失败"
		} else {
			dataMap["code"] = 0
			dataMap["data"] = list
			dataMap["message"] = "查询数据成功"
		}
		db.Close()
	}
	c.JSON(http.StatusOK, dataMap)
}

func AddCoursePlane(c *gin.Context) {
	courseid := c.PostForm("courseid")
	stcount := c.PostForm("stcount")
	courseaddress := c.PostForm("courseaddress")
	coursetecher := c.PostForm("coursetecher")
	coursephone := c.PostForm("coursephone")
	courseemail := c.PostForm("courseemail")
	coursetarget := c.PostForm("coursetarget")
	coursetime := c.PostForm("coursetime")
	dataMap := make(map[string]interface{})
	if len(courseid) == 0 || len(stcount) == 0 || len(courseaddress) == 0 || len(coursetime) == 0 {
		dataMap["code"] = 2
		dataMap["message"] = "请填写完整必填项"
	} else {
		db, err := utils.OpenGinDB()
		if err != nil {
			dataMap["code"] = 3
			dataMap["message"] = "打开数据库失败"
		} else {
			cout, cerr := db.Table("course").Where("courseid", courseid).Count("*")
			if cerr != nil {
				dataMap["code"] = 4
				dataMap["errinfo"] = cerr
				dataMap["message"] = "查询课程失败"
			} else {
				if cout == 0 {
					dataMap["code"] = 4
					dataMap["message"] = "没有查询到相应的课程"
				} else {
					if len(courseaddress) == 0 {
						courseaddress = ""
					}
					if len(coursetecher) == 0 {
						coursetecher = ""
					}
					if len(coursephone) == 0 {
						coursephone = ""
					}
					if len(courseemail) == 0 {
						courseemail = ""
					}
					if len(coursetarget) == 0 {
						coursetarget = ""
					}
					params := make(map[string]interface{})
					params["chooseid"] = "CHO" + utils.CreateStudentid()
					params["courseid"] = courseid
					intStcount, sterr := strconv.ParseInt(stcount, 10, 64)
					if sterr != nil {
						intStcount = 0
					}
					params["stcount"] = intStcount
					params["courseaddress"] = courseaddress
					params["coursetecher"] = coursetecher
					params["coursephone"] = coursephone
					params["courseemail"] = courseemail
					params["coursetarget"] = coursetarget
					params["hascount"] = 0
					params["leftcount"] = 0
					params["coursetime"] = coursetime
					_, inerr := db.Table("course_choose").Data(params).Insert()
					if inerr != nil {
						dataMap["code"] = 3
						dataMap["errinfo"] = inerr
						dataMap["message"] = "出入数据库失败"
					} else {
						dataMap["code"] = 0
						dataMap["message"] = "添加课程计划成功"
					}
				}
			}
			db.Close()
		}
	}
	c.JSON(http.StatusOK, dataMap)
}

func GetCoursePlanList(c *gin.Context) {
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
		dataMap["message"] = "打开数据库失败"
	} else {
		intlimit, limiterr := strconv.Atoi(limitstr)
		intoffset, offseterr := strconv.Atoi(offsetstr)
		if limiterr != nil {
			intlimit = 20
		}
		if offseterr != nil {
			intoffset = 20
		}
		list, cerr := db.Table("course_choose").Fields("chooseid,courseaddress,courseemail,courseid,coursephone,coursetarget,coursetecher,coursetime,hascount,leftcount,stcount").Limit(intlimit).Offset(intoffset).Get()
		if cerr != nil {
			dataMap["code"] = 2
			dataMap["errinfo"] = cerr
			dataMap["message"] = "查询数据库失败"
		} else {
			dataMap["code"] = 0
			dataMap["data"] = list
			dataMap["message"] = "查询数据成功"
		}
		db.Close()
	}
	c.JSON(http.StatusOK, dataMap)
}

func AddSelectCoursePlan(c *gin.Context) {
	dataMap := make(map[string]interface{})
	chooseid := c.PostForm("chooseid")
	studentid := c.PostForm("studentid")
	if len(chooseid) == 0 || len(studentid) == 0 {
		dataMap["code"] = 2
		dataMap["message"] = "请填写完整必填项"
	} else {
		db, openerr := utils.OpenGinDB()
		if openerr != nil {
			dataMap["code"] = 4
			dataMap["errinfo"] = openerr
			dataMap["message"] = "打开数据库失败"
		} else {
			//首先看学生存不存在
			//然后看课程存不存在
			//再判断
			//最后看课程是否选满
			//添加选择的课程计划id和学生id
			//查询学生
			count, cerr := db.Table("studentInfo").Where("studentid", studentid).Count("*")
			if cerr != nil {
				dataMap["code"] = 2
				dataMap["errinfo"] = cerr
				dataMap["message"] = "查询学生数据库失败"
			} else {
				if count > 0 {
					//查询课程计划是否存在
					plan, planerr := db.Table("course_choose").Where("chooseid", chooseid).First()
					if planerr != nil {
						dataMap["code"] = 2
						dataMap["errinfo"] = cerr
						dataMap["message"] = "查询课程计划库失败"
					} else {
						if len(plan) == 0 {
							dataMap["code"] = 2
							dataMap["errinfo"] = cerr
							dataMap["message"] = "没有查询到相应的课程计划"
						} else {
							//判断这个学生的这个课程计划是否添加
							stplancount, stplanerr := db.Table("student_course_choose").Where("studentid", studentid).Where("chooseid", chooseid).Count("*")
							if stplanerr != nil {
								dataMap["code"] = 2
								dataMap["errinfo"] = cerr
								dataMap["message"] = "查询课程计划库失败"
							} else {
								if stplancount > 0 {
									dataMap["code"] = 6
									dataMap["message"] = "你已经选择该课程"
								} else {
									allCount, allok := plan["stcount"]
									hasCount, hasok := plan["hascount"]
									leftCount, leftok := plan["leftcount"]
									if hasok == true && leftok == true && allok == true {
										intall := allCount.(int64)
										inthas := hasCount.(int64)
										intleft := leftCount.(int64)
										if (inthas + 1) > intall {
											dataMap["code"] = 5
											dataMap["message"] = "对不起，改课程计划已经选满"
										} else {
											//需要放在事务里面
											db.Begin()
											stplanParams := make(map[string]interface{})
											stplanParams["studentid"] = studentid
											stplanParams["chooseid"] = chooseid
											_, insrterr := db.Table("student_course_choose").Data(stplanParams).Insert()
											if insrterr != nil {
												dataMap["code"] = 7
												dataMap["errinfo"] = insrterr
												dataMap["message"] = "选课成功,插入数据库失败"
												db.Rollback()
											} else {
												inthas = inthas + 1
												intleft = intall - inthas
												updateParams := make(map[string]interface{})
												updateParams["hascount"] = inthas
												updateParams["leftcount"] = intleft
												_, updaterr := db.Table("course_choose").Data(updateParams).Where("chooseid", chooseid).Update()
												if updaterr != nil {
													dataMap["code"] = 7
													dataMap["errinfo"] = updaterr
													dataMap["message"] = "选课成功,更新数据库失败"
													db.Rollback()
												} else {
													dataMap["code"] = 0
													dataMap["message"] = "选课成功"
												}
											}
											db.Commit()
										}
									} else {
										dataMap["code"] = 4
										dataMap["message"] = "没有查询到课程计划的数据"
									}
								}
							}
						}
					}

				} else {
					dataMap["code"] = 3
					dataMap["message"] = "没有查询到学生信息"
				}
			}
		}
		db.Close()
	}
	c.JSON(http.StatusOK, dataMap)
}

//获取选取这个课程计划的学生信息
func GetInfoListByChooseid(c *gin.Context) {
	dataMap := make(map[string]interface{})
	chooseid := c.PostForm("chooseid")
	limitstr := c.PostForm("limit")
	offsetstr := c.PostForm("offset")
	if len(limitstr) == 0 || len(offsetstr) == 0 {
		limitstr = "20"
		offsetstr = "0"
	}
	db, err := utils.OpenGinDB()
	if err != nil {
		dataMap["code"] = 4
		dataMap["message"] = "打开数据库失败"
	} else {
		intlimit, limiterr := strconv.Atoi(limitstr)
		intoffset, offseterr := strconv.Atoi(offsetstr)
		if limiterr != nil {
			intlimit = 20
		}
		if offseterr != nil {
			intoffset = 0
		}
		if len(chooseid) == 0 {
			list, cherr := db.Table("studentInfo").Fields("studentInfo.name,studentInfo.studentid,studentInfo.sex,studentInfo.phone,studentInfo.guardian,studentInfo.grade,studentInfo.class,studentInfo.age,studentInfo.address").Join("student_course_choose", "studentInfo.studentid", "=", "student_course_choose.studentid").Limit(intlimit).Offset(intoffset).Get()
			if cherr != nil {
				dataMap["code"] = 2
				dataMap["errinfo"] = cherr
				dataMap["message"] = "查询数据失败"
				lastsql := db.LastSql()
				fmt.Println(lastsql)
			} else {
				dataMap["code"] = 0
				dataMap["data"] = list
				dataMap["message"] = "查询数据成功"
			}
		} else {
			list, cherr := db.Table("studentInfo").Fields("studentInfo.name,studentInfo.studentid,studentInfo.sex,studentInfo.phone,studentInfo.guardian,studentInfo.grade,studentInfo.class,studentInfo.age,studentInfo.address").Join("student_course_choose", "studentInfo.studentid", "=", "student_course_choose.studentid").Where("chooseid", chooseid).Limit(intlimit).Offset(intoffset).Get()
			if cherr != nil {
				dataMap["code"] = 2
				dataMap["errinfo"] = cherr
				dataMap["message"] = "查询数据失败"
			} else {
				dataMap["code"] = 0
				dataMap["data"] = list
				dataMap["message"] = "查询数据成功"
			}
		}
		db.Close()

	}
	c.JSON(http.StatusOK, dataMap)
}

//获取当前学生的选课信息
func GetInfoListByStudentid(c *gin.Context) {
	dataMap := make(map[string]interface{})
	studentid := c.PostForm("studentid")
	limitstr := c.PostForm("limit")
	offsetstr := c.PostForm("offset")
	if len(limitstr) == 0 || len(offsetstr) == 0 {
		limitstr = "20"
		offsetstr = "0"
	}
	db, err := utils.OpenGinDB()
	if err != nil {
		dataMap["code"] = 4
		dataMap["message"] = "打开数据库失败"
	} else {
		intlimit, limiterr := strconv.Atoi(limitstr)
		intoffset, offseterr := strconv.Atoi(offsetstr)
		if limiterr != nil {
			intlimit = 20
		}
		if offseterr != nil {
			intoffset = 20
		}
		if len(studentid) == 0 {
			list, cherr := db.Query("SELECT T0.chooseid,T0.courseaddress,T0.courseemail,T0.courseid,T0.coursephone,T0.coursetarget,T0.coursetecher,T0.courseTime,T0.hascount,T0.leftcount,T0.stcount,T2.courseName,T2.isrequired,T2.profession from course_choose T0 INNER JOIN student_course_choose T1 ON T0.chooseid=T1.chooseid INNER JOIN course T2 ON T0.courseid=T2.courseid  LIMIT ? OFFSET ?", intlimit, intoffset)
			if cherr != nil {
				dataMap["code"] = 2
				dataMap["errinfo"] = cherr
				dataMap["message"] = "查询数据库失败"
			} else {
				dataMap["code"] = 0
				dataMap["data"] = list
				dataMap["message"] = "获取数据成功"
			}
		} else {
			list, cherr := db.Query("SELECT T0.chooseid,T0.courseaddress,T0.courseemail,T0.courseid,T0.coursephone,T0.coursetarget,T0.coursetecher,T0.courseTime,T0.hascount,T0.leftcount,T0.stcount,T2.courseName,T2.isrequired,T2.profession from course_choose T0 INNER JOIN student_course_choose T1 ON T0.chooseid=T1.chooseid INNER JOIN course T2 ON T0.courseid=T2.courseid WHERE T1.studentid=？ LIMIT ? OFFSET ?", studentid, intlimit, intoffset)
			if cherr != nil {
				dataMap["code"] = 2
				dataMap["errinfo"] = cherr
				dataMap["message"] = "查询数据库失败"
			} else {
				dataMap["code"] = 0
				dataMap["data"] = list
				dataMap["message"] = "获取数据成功"
			}
		}
		db.Close()
	}
	c.JSON(http.StatusOK, dataMap)
}
