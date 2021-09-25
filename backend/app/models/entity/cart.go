package entity

import "time"

type Cart struct{
	ID int `gorm:"primaryKey;autoIncrement:true"`
	CreatedAt time.Time
	Status int
	TotalPrice float64

	CartDetails [] CartDetail `gorm:"foreignKey:CartId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

}

func (Cart) TableName() string {
    return "db_cart"
}