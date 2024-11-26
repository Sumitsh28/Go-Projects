package models

import (
	"go_learn/my_sql/pkg/configs"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Car struct{
	gorm.Model

	Name string `gorm:""json:"name"`
	Manufacturer string `json:"manufacturer"`
	Year string `json:"year"`
}

func init(){
	configs.Connect()
	db = configs.GetDB()
	db.AutoMigrate(&Car{})
}

func (c *Car) CreateCar() *Car {

	db.NewRecord(c)
	db.Create(&c)

	return c
}

func GetAllCars() []Car{

	var Cars []Car
	db.Find(&Cars)
	return Cars
}

func GetCarById(Id int64) (*Car, *gorm.DB){

	var getCar Car
	db := db.Where("ID = ?", Id).Find(&getCar)

	return &getCar, db
}

func DeleteCar(ID int64) Car{
	var car Car
	db.Where("ID=?", ID).Delete(car)
	return car
}