package entity

type Category struct{
	ID int `gorm:"primaryKey;autoIncrement:true"`
	UserId int 
	Cost float64 
	Status int

	Products []Product `gorm:"foreignKey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Category) TableName() string {
    return "db_category"
}