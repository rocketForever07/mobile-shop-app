package entity

type Brand struct {

	ID        int `gorm:"primaryKey;autoIncrement:true"`
	BrandName string
	Products []Product `gorm:"foreignKey:BrandId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

					  
}

func (Brand) TableName() string {
    return "db_brand"
}