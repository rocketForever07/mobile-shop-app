package entity

import(
	"gorm.io/gorm"
)

type Product struct{
	ID 			int `gorm:primaryKey` 
	Code		string `gorm:unique`
	Name 		string `gorm:unique`
	Price 		float64
	Quantity 	int
	Description string
	Image		string
	Color		string
	Memory		string
	Active 		int
	CategoryId 	int
	BrandId		int
	OrderDetails [] OrderDetail `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CartDetails  [] CartDetail 	`gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}


func (Product) TableName() string {
    return "db_product"
}


type ProductStorage struct{
	Storage *gorm.DB
}


func (ps ProductStorage) GetList() ([]Product,error){
	var products []Product

	result := ps.Storage.Where("active = ?",1).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (ps ProductStorage) GetById(id int) (Product,error){
	var product Product

	result := ps.Storage.First(&product,id)
	if result.Error != nil {
		return Product{}, result.Error
	}
	return product, nil
}

func (ps ProductStorage) GetByCode(code string) ([]Product,error){

	var products []Product

	result := ps.Storage.Where("code = ?", code).Find(&products)
	if result.Error != nil{
		return nil,result.Error
	}

	return products, nil
	
}

func (ps ProductStorage) Insert(product Product) (Product,error){
	result := ps.Storage.Omit("OrderDetails","CartDetails").Create(&product)

	if result.Error!=nil{
		return Product{},result.Error
	}

	return product,nil
}

func (ps ProductStorage) Update(product Product) (Product,error){
	result := ps.Storage.Save(&product)

	if result.Error!=nil{
		return Product{},result.Error
	}

	return product,nil
}

func (ps ProductStorage) Delete(id int) (Product,error){
	var product Product
	ps.Storage.First(&product,id)
	product.Active=0
	
	result := ps.Storage.Save(&product)

	if result.Error!=nil{
		return Product{},result.Error
	}

	return product,nil

}

func (ps ProductStorage) GetByNameContaining(nameQuerry string) []Product{
	return  []Product{}
}



