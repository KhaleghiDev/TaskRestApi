package main

import (
	"RestApi/config"
	"RestApi/model"
	"RestApi/rooter"
	"fmt"
)

func main() {

	//TODO :
	// 1 - config server and database
	// 2-  start router and cheack postman
	// 2-1 start swager
	// 3 - user controller and auth register ,login
	// 4 - response json and test register ,login
	// 5 - start tickeate  and model and controller
	// 6 - list controller
	// 7-  cheack functino and matiol
	// 8-  tester
	fmt.Println("start projeact")
	dbconf := config.ConfDb{
		Hose:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "",
		Db:       "test",
	}
	db := config.ConnMysql(dbconf)
	db.AutoMigrate(new(model.User))
	db.AutoMigrate(new(model.Ticket))
	db.AutoMigrate(new(model.Response))
	addr := "localhost:8000"
	rooter.Rooter(addr, db)
}
