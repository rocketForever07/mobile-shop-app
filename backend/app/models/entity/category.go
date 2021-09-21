package entity

type Category struct{
	ID int `gorm:"primaryKey;autoIncrement:true"`
	Name string

	Products []Product `gorm:"foreignKey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Category) TableName() string {
    return "db_category"
}


func (Category) GetList() []Category{
	return  []Category{}
}

func (Category) GetById(id int) Category{
	return Category{}
}

func (Category) Insert(category Category){

}

func (Category) Update(category Category,id int){

}

func (Category) Delete(id int){

}

