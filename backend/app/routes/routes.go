package routes

import(
	"app/handler"
	"github.com/gorilla/mux"
	db "app/database"
	"app/models/entity"
	"log"
	"net/http"


)

func Init(){

	/*Declare handler model*/
	productHandler := handler.ProductHandler{entity.ProductStorage{db.DB}}


	/*Declare router*/
	//main route
	router := mux.NewRouter().StrictSlash(true)
	
	//product route
	productRoute := router.PathPrefix("/api/product").Subrouter()


	/**/
	//handle for product route
	productRoute.HandleFunc("",productHandler.GetAll).Methods("GET")
	productRoute.HandleFunc("/{code}",productHandler.GetDetail).Methods("GET")
	productRoute.HandleFunc("",productHandler.Create).Methods("POST")
	productRoute.HandleFunc("/{id}",productHandler.Update).Methods("PUT")
	productRoute.HandleFunc("/{id}",productHandler.Delete).Methods("DELETE")


	log.Println("listen server on port 10000")
	log.Fatal(http.ListenAndServe(":10000",router))
	http.ListenAndServe(":10000",router)
	
	

}