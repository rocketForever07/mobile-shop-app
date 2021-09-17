package entity

type User struct{
	ID 			int `gorm:primaryKey`
	Username 	string 
	Password 	string
	Fullname 	string 
	Age 		int
	Address 	string
	UserImage 	string
	
	Orders []Order `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Carts  []Cart  `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

}

func (User) TableName() string {
    return "db_user"
}