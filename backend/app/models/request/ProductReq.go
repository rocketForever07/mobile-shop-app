package request

import(
	"app/models/entity"
)

type ProductReq struct{
	Code		string `json:"code"`
	Name 		string `json:"name"`
	Price 		float64	`json:"price"`
	Quantity 	int	`json:"quantity"`
	Color		string `json:"color"`
	Memory		string`json:"memory"`
	Description string `json:"description"`
	CategoryId 	int	`json:"categoryId"`
	BrandId		int	`json:"brandId"`
	Image		int `json:"image"`
}

//convert from req to entity
func (pr ProductReq) Convert() entity.Product{

	productReq := entity.Product{
		
		pr.Code,
		pr.Name,
		pr.Price,
		pr.Quantity,
		pr.Description,
		pr.Color,
		pr.Memory,
		pr.CategoryId,
		pr.BrandId,
		pr.Image,

	}

	return productReq

}
