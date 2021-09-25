package entity

import(
	"gorm.io/gorm"
)

type OrderDetail struct {
	ID        int `gorm:"primaryKey;autoIncrement:true"`
	ProductId int 
	OrderId   int    
	Quantity  int	
	Price 	  float64
					  
}

func (OrderDetail) TableName() string {
    return "db_order_detail"
}

type OrderDetailStorage struct{
	Storage *gorm.DB
}


func (ods OrderDetailStorage) GetById(id int) (OrderDetail,error){
	var orderDetail OrderDetail

	result := ods.Storage.First(&orderDetail,id)
	if result.Error != nil {
		return OrderDetail{}, result.Error
	}
	return orderDetail, nil
}


func (ods OrderDetailStorage) Insert(orderDetail OrderDetail) (OrderDetail,error){
	result := ods.Storage.Create(&orderDetail)

	if result.Error!=nil{
		return OrderDetail{},result.Error
	}

	return orderDetail,nil
}

func (ods OrderDetailStorage) Update(orderDetail OrderDetail) (OrderDetail,error){
	result := ods.Storage.Save(&orderDetail)

	if result.Error!=nil{
		return OrderDetail{},result.Error
	}

	return orderDetail,nil
}

func (ods OrderDetailStorage) Delete(id int) (error){
	
	result := ods.Storage.Delete(id)

	if result.Error!=nil{
		return result.Error
	}

	return nil

}