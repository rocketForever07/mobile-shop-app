package entity

import(
	"gorm.io/gorm"
	"time"
)

type Order struct{
	ID 				int `gorm:"primaryKey;autoIncrement:true"`
	UserId 			int 
	TotalPrice 		float64 
	CustomerName 	string
	CustomerPhone 	string
	CustomerEmail 	string
	DeliveryAddress	string
	Description 	string
	OrderDate       time.Time
	Status 			int

	OrderDetails []OrderDetail `gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Order) TableName() string {
    return "db_order"
}


type OrderStorage struct{
	Storage *gorm.DB
}

func (os OrderStorage) GetList() ([]Order,error){
	var products []Order

	result := os.Storage.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (os OrderStorage) GetById(id int) (Order,error){
	var order Order

	result := os.Storage.First(&order,id)
	if result.Error != nil {
		return Order{}, result.Error
	}
	return order, nil
}


func (os OrderStorage) Insert(order Order) (Order,error){
	result := os.Storage.Omit("OrderDetails").Create(&order)

	if result.Error!=nil{
		return Order{},result.Error
	}

	return order,nil
}

func (os OrderStorage) Update(order Order) (Order,error){
	result := os.Storage.Save(&order)

	if result.Error!=nil{
		return Order{},result.Error
	}

	return order,nil
}

func (os OrderStorage) Delete(id int) (error){
		
	result := os.Storage.Delete(id)

	if result.Error!=nil{
		return result.Error
	}

	return nil

}




