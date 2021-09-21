package repository

import(
	"app/models/entity"
	db "app/database"
	"gorm.io/gorm" 
)

type ProductRepo struct{
	repo *gorm.DB
}


func (pr ProductRepo) GetList() ([]entity.Product,error){
	var products []entity.Product

	result := db.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (ProductRepo) GetById(id int) (entity.Product,error){
	var product entity.Product

	result := db.DB.First(&product,id)
	if result.Error != nil {
		return entity.Product{}, result.Error
	}
	return product, nil
}

func (ProductRepo) Insert(product entity.Product) (entity.Product,error){
	result := db.DB.Omit("OrderDetails","CartDetails").Create(&product)

	if result.Error!=nil{
		return entity.Product{},result.Error
	}

	return product,nil
}

func (ProductRepo) Update(product entity.Product,id int) (entity.Product,error){
	result := db.DB.Save(&product)

	if result.Error!=nil{
		return entity.Product{},result.Error
	}

	return product,nil
}

func (ProductRepo) Delete(id int) error{ 
	result := db.DB.Delete(entity.Product{},id)

	return result.Error
}

func (ProductRepo) GetByNameContaining(nameQuerry string) []entity.Product{
	return  []entity.Product{}
}