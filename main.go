package main

import(
	"app/route"
	"app/database"
	"app/model"
)

func main(){

	Conn := database.GetConnection()
	// Migrate the schema
	Conn.DB.AutoMigrate(&model.User{})

	route.HandleRoute()
}