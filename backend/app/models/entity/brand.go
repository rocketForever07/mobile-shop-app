package entity

import(
	"gorm.io/gorm"
)

type Brand struct {

	ID        int `gorm:"primaryKey;autoIncrement:true"`
	BrandName string
	Products []Product `gorm:"foreignKey:BrandId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

					  
}

func (Brand) TableName() string {
    return "db_brand"
}

type BrandStorage struct{
	Storage *gorm.DB
}


func (bs BrandStorage) GetNameById(id int) string{
	var brandEntity Brand

	bs.Storage.First(&brandEntity,id)
	return brandEntity.BrandName
}