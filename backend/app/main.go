package main

import(
	db "app/database"
	"app/routes"
	"app/helper/scheduler"
)

func main(){

	//change me!
	config := db.Config{
		Server:"localhost:3306",
		Username:"root",
		Password:"dtkhoi07",
		DB:"mobile_shop_app",
	}

	db.InitDB(config)
	routes.Init()
	scheduler.start(db.DB)
	
}