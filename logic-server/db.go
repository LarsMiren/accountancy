package main

import (
	"fmt"

	"github.com/golang/glog"

	"github.com/LarsMiren/accountancy/env"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID            uint   `gorm:"primary_key; AUTO_INCREMENT"`
	Username      string `gorm:"column:username"`
	Email         string `gorm:"column:email;unique_index"`
	PersonalData  string `gorm:"column:personal_data;size:1024"`
	Image         []byte `gorm:"column:image"`
	PasswordHash  string `gorm:"column:password;not null"`
	Subscriptions []User
}

type Product struct {
	ID          uint64 `gorm:"primary_key",sql:"AUTO_INCREMENT"`
	Name        string
	Image       []byte
	Description string
	Cost        float32
	Supplier    User
	Type        ProductType
}

type ProductType struct {
	ID   uint64 `gorm:"primary_key, AUTO_INCREMENT"`
	Name string
}

var (
	db *gorm.DB
)

func connectToDB() *gorm.DB {
	dbOpts := env.GetDB("postgres")
	opts := "host=" + env.GetAddr("postgres") +
		" port=" + env.GetPort("postgres") +
		" user=" + dbOpts.User +
		" password=" + dbOpts.Password +
		" dbname=" + dbOpts.DBName
	fmt.Println("Connecting to DB server...")
	db, err := gorm.Open("postgres", opts)
	check(err)
	if !dbOpts.WithTables {
		fmt.Println("Creating tables...")
		createTables()
	}
	return db
}

func createTables() {
	models := []interface{}{&User{}, &ProductType{}, &Product{}}
	fmt.Println(models[1])
	fmt.Println("******")
	db = db.DropTableIfExists(models)
	fmt.Println("tables droped")
	db = db.CreateTable(models)
	if err := db.Error; err != nil {
		glog.Fatalln(err)
		panic(err)
	}
}
