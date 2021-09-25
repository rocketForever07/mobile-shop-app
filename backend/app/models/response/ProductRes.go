package response

import(
	"app/models/entity"
	db "app/database"
)



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
	BrandName	string `json:"brandName"`
	Image		string `json:"image"`
}


func  ProductMapper(e entity.Product) ProductRes{
	productRes := ProductRes{e.ID,e.Code,e.Name,e.Price,e.Quantity,e.Description,e.Color,e.Memory,e.CategoryId,e.BrandId,"",e.Image}

	//declare storage for query
	brandStorage := entity.BrandStorage{db.DB}

	//set brand name
	productRes.BrandName=brandStorage.GetNameById(e.BrandId)

	return productRes
}