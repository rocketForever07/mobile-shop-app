package main

import(
	db "app/database"
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

	

}