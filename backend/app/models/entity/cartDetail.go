package entity

import(
	"gorm.io/gorm"
)

type CartDetail struct{

	ID int `gorm:"primaryKey;autoIncrement:true"`
	CartId int
	ProductId int
	Quantity int
	Price float64
}


func (CartDetail) TableName() string {
    return "db_cart_detail"
}


type CartDetailStorage struct{
	Storage *gorm.DB
}


func (cds CartDetailStorage) GetById(id int) (CartDetail,error){
	var cartDetail CartDetail

	result := cds.Storage.First(&cartDetail,id)
	if result.Error != nil {
		return CartDetail{}, result.Error
	}
	return cartDetail, nil
}


func (cds CartDetailStorage) Insert(cartDetail CartDetail) (CartDetail,error){
	result := cds.Storage.Create(&cartDetail)

	if result.Error!=nil{
		return CartDetail{},result.Error
	}

	return cartDetail,nil
}

func (cds CartDetailStorage) Update(cartDetail CartDetail) (CartDetail,error){
	result := cds.Storage.Save(&cartDetail)

	if result.Error!=nil{
		return CartDetail{},result.Error
	}

	return cartDetail,nil
}

func (cds CartDetailStorage) Delete(id int) (error){
	
	result := cds.Storage.Delete(id)

	if result.Error!=nil{
		return result.Error
	}

	return nil

}

