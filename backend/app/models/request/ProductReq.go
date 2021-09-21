package models

type ProductReq struct{
	Code		string `json:"code"`
	Name 		string `json:"name"`
	Price 		float64	`json:"price"`
	Quantity 	int	`json:"quantity"`
	Description string `json:"description"`
	CategoryId 	int	`json:"categoryId"`
	BrandId		int	`json:"brandId"`
}