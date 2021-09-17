package entity

type Product struct{
	ID 			int `gorm:primaryKey` 
	Code		string
	Name 		string
	Price 		float64
	Quantity 	int
	Description string
	Active 		int
	CategoryId 	int
	BrandId		int
	OrderDetails [] OrderDetail `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CartDetails [] CartDetail 	`gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}


func (Product) TableName() string {
    return "db_product"
}


