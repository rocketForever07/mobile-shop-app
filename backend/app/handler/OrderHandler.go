package handler

import (
	"log"
	"encoding/json"
	"fmt"
	"strconv"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"app/models/entity"
	res "app/models/response"

)


type OrderHandler struct{
	Repo entity.OrderStorage
}


func (ph OrderHandler) GetDetail(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id,_:= strconv.Atoi(vars["id"])

	var order entity.Order
	var err error

	order,err = ph.Repo.GetById(id)
	
	fmt.Println(order)
	if err != nil{
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}

	//mapper entity to response
	var ordersRes res.OrderRes
	
	ordersRes=res.OrderMapper(order)
	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ordersRes)

}

func (ph OrderHandler) Create(w http.ResponseWriter, r *http.Request){

	reqBody,err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
	}

	var order entity.Order
	//convert json to obj
	json.Unmarshal(reqBody,&order)
	fmt.Println(order)

	//import to db
	newProduct,err := ph.Repo.Insert(order)

	if err != nil{
		log.Println(err)	
		w.WriteHeader(http.StatusNoContent)
	}
	//
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	//response this payment obj(json)
	json.NewEncoder(w).Encode(newProduct)

}

//only cancel order
func (ph OrderHandler) Update(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id,_ := strconv.Atoi(vars["id"])

	requestBody, _ := ioutil.ReadAll(r.Body)
	var order entity.Order
	json.Unmarshal(requestBody, &order)
	order.ID = id
	log.Println(order)

	updateProduct,err := ph.Repo.Update(order)
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateProduct)

}

