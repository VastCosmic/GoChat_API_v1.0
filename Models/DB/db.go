package DB

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect 连接数据库
func Connect() *gorm.DB {
	// 使用os.ReadFile函数读取db_config.json文件的内容，得到一个字节切片
	dbConfigJSON, err := os.ReadFile("Models/DB/dbConfig.json")
	if err != nil {
		log.Fatal(err)
	}

	// 声明一个DBConfig变量，用于接收转换后的数据
	var dbConfig DBConfig

	// 使用json.Unmarshal函数将dbConfigJSON转换成DBConfig结构体，并赋值给dbConfig变量
	err = json.Unmarshal(dbConfigJSON, &dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 使用dbConfig中的字段来拼接连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Dbname, dbConfig.Timeout)
	//println(dsn)

	// 使用mysql驱动和连接字符串打开一个数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 连接成功
	//fmt.Println(db)
	//fmt.Println("mysql is ready")
	return db
}
