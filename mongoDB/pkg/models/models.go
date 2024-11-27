package models

import (
	"context"
	"go_learn/mongoDB/pkg/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Manufacturer string             `json:"manufacturer" bson:"manufacturer"`
	Year         string             `json:"year" bson:"year"`
}

var collection = config.GetCollection("cars")

func CreateCar(car Car) (Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	car.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, car)
	return car, err
}

func GetAllCars() []Car {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var cars []Car
	for cursor.Next(ctx) {
		var car Car
		cursor.Decode(&car)
		cars = append(cars, car)
	}
	return cars
}

func GetCarById(carId string) (Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(carId)
	var car Car
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&car)
	return car, err
}

func UpdateCar(carId string, updateCar Car) (Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(carId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateCar}

	_, err := collection.UpdateOne(ctx, filter, update)
	return updateCar, err
}

func DeleteCar(carId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(carId)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
