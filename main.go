package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int
	Email    string
	Password string
	Profile  Profile `gorm:"foreignKey:User_id"`
}

type Profile struct {
	ID      int
	Fname   string
	Lname   string
	Address string
	Phone   string
	User_id int
}

func main() {
	dsn := "host=localhost user=postgres password=admin dbname=migrate_ngc10 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("unable to run the db: ", err)
	}

	res := db.AutoMigrate(&User{}, &Profile{})
	if res != nil {
		log.Println(err)
	}
}
