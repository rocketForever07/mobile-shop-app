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


type ProductHandler struct{
	Repo entity.ProductStorage
}


func (ph ProductHandler) GetAll(w http.ResponseWriter, r *http.Request){

	var products []entity.Product
	var err error

	products,err = ph.Repo.GetList()

	if err != nil{
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}else{ 
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		//mapper entity to response
		var productsRes []res.ProductRes
		for _,v := range products{
			productsRes=append(productsRes, res.ProductMapper(v))
		}

		json.NewEncoder(w).Encode(productsRes)
	}

}

func (ph ProductHandler) GetDetail(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	code:= vars["code"]

	var products []entity.Product
	var err error

	products,err = ph.Repo.GetByCode(code)
	
	fmt.Println(products)
	if err != nil{
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}

	//mapper entity to response
	var productsRes []res.ProductRes
	for i:=0;i<len(products);i++{
		productsRes=append(productsRes, res.ProductMapper(products[i]))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productsRes)

}

func (ph ProductHandler) Create(w http.ResponseWriter, r *http.Request){

	reqBody,err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Println(err)
	}

	var product entity.Product
	//convert json to obj
	json.Unmarshal(reqBody,&product)
	fmt.Println(product)

	//import to db
	newProduct,err := ph.Repo.Insert(product)

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

func (ph ProductHandler) Update(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id,_ := strconv.Atoi(vars["id"])

	requestBody, _ := ioutil.ReadAll(r.Body)
	var product entity.Product
	json.Unmarshal(requestBody, &product)
	product.ID = id
	log.Println(product)

	updateProduct,err := ph.Repo.Update(product)
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateProduct)

}

func  (ph ProductHandler) Delete(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id,_ := strconv.Atoi(vars["id"])

	deleteProduct,err := ph.Repo.Delete(id)
	if err!=nil{
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deleteProduct)
	w.WriteHeader(http.StatusOK)
}
