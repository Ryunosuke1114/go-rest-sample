package main
import (
	"log"
	"xorm.io/xorm"
)

func DBInit() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql" , "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil{
		log.Fatal(err)
	}

	engine.ShowSQL(true)

	exist, err := engine.IsTableExist("users")
	if err != nil {
		log.Fatal(err)
	}

	if !exist{
		engine.CreateTables(&model.Users{})
	}
	return engine
}