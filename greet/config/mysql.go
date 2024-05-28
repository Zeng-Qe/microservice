package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 创建一个全局变量来存储数据库连接池
//var db *sql.DB

func DBInit() {
	// 配置数据库连接字符串
	//addr := viper.GetString("db.addr")
	//port := viper.GetInt("db.port")
	user := viper.GetString("db.user")
	pass := viper.GetString("db.password")
	dbname := viper.GetString("db.dbname")
	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, pass, dbname)
	// 连接到数据库
	var err error
	//db, err = sql.Open("mysql", connStr)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.DB()
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	//r, err1 := sqlDB.Exec("insert into users(name, phone_number, email,password)values(?, ?, ?,?)", "aaa", "123456", "aaa", "123456")
	//if err1 != nil {
	//	fmt.Println("数据插入异常, ", err1)
	//	return
	//}
	//_, err2 := r.LastInsertId()
	//if err2 != nil {
	//	fmt.Println("获取id异常:, ", err2)
	//	return
	//}
	fmt.Println("Database connection successful.")

	// 使用连接池执行查询
	db.Exec("SELECT * FROM users")
}
