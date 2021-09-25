package database

import(
	"gorm.io/gorm"
	"log"
	entity "app/models/entity"
	"gorm.io/driver/mysql"

)

var DB *gorm.DB
var err error
func InitDB(config Config){

	dsn := GetConnectionString(config)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database %s", err)
	}

	DB.AutoMigrate(&entity.User{},
		&entity.Category{},  
		&entity.Brand{}, 
		&entity.Product{}, 
		&entity.Payment{}, 
		&entity.Order{}, 
		&entity.OrderDetail{}, 
		&entity.Cart{}, 
		&entity.CartDetail{})
	if err != nil {
		log.Fatal("failed to load migrations")
	}

	log.Println("successful migrate")

}
