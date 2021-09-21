package models

type ProductRes struct{
	ID 			int `json:"id"`
	Code		string `json:"code"`
	Name 		string `json:"name"`
	Price 		float64	`json:"price"`
	Quantity 	int	`json:"quantity"`
	Description string `json:"description"`
	Color 		string `json:"color"` 
	Memory		string `json:"memory"`
	CategoryId 	int	`json:"categoryId"`
	BrandId		int	`json:"brandId"`
	BrandName	int `json:"brandName"`
}
