package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //import DB driver
	"github.com/gohouse/gorose"        //import Gorose
	"math/rand"
	"strconv"
	"time"
)

var DbConfig = map[string]interface{}{
	// Default database configuration
	"Default": "mysql_dev",
	// (Connection pool) Max open connections, default value 0 means unlimit.
	"SetMaxOpenConns": 300,
	// (Connection pool) Max idle connections, default value is 1.
	"SetMaxIdleConns": 10,

	// Define the database configuration character "mysql_dev".
	"Connections": map[string]map[string]string{
		"mysql_dev": map[string]string{
			"host":     "127.0.0.1",
			"username": "root",
			"password": "medex",
			"port":     "3306",
			"database": "medex",
			"charset":  "utf8",
			"protocol": "tcp",
			"prefix":   "",      // Table prefix
			"driver":   "mysql", // Database driver(mysql,sqlite,postgres,oracle,mssql)
		},
	},
}

func OpenGinDB() (con *gorose.Connection, err error) {
	fmt.Println("开始打开数据库")
	db, err := gorose.Open(DbConfig)
	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	} else {
		fmt.Println("连接数据库成功")
	}
	// close DB
	// res, err := db.Table("studentInfo").First()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(res)
	// defer db.Close()
	return &db, err
}

func CreateStudentid() string {
	now := time.Now()
	//timeStamp := now.Unix()
	timeStamp := now.UnixNano() / 1000000
	r := rand.New(rand.NewSource(now.UnixNano()))
	stampStr := strconv.FormatInt(timeStamp, 10)
	orderId := fmt.Sprintf("%s%d", stampStr, r.Intn(100))
	return orderId
}
