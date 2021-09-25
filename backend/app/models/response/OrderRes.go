package response

import(
	"app/models/entity"
	"time"
)

type OrderRes struct{
	ID 				int `gorm:"primaryKey;autoIncrement:true"`
	TotalPrice 		float64 
	CustomerName 	string
	CustomerPhone 	string
	CustomerEmail 	string
	DeliveryAddress	string
	Description 	string
	OrderDate       time.Time
	Status 			int
}



func OrderMapper(e entity.Order) OrderRes{
	
	orderRes := OrderRes{
		e.ID,
		e.TotalPrice,
		e.CustomerName,
		e.CustomerPhone,
		e.CustomerEmail,
		e.DeliveryAddress,
		e.Description,
		e.OrderDate,
		e.Status,
	}

	return orderRes
}