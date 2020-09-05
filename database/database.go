package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"fmt"
)

type Connection struct { 
	DB *gorm.DB
}

func (conn *Connection) OpenConnection( ) {
	if conn.DB == nil {
		fmt.Println("Reconnect")
		dsn := "root:Fonseca@12@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil{
			log.Fatal(err)
		}

		conn.DB = db;
	}

}
