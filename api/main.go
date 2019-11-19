package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Choice struct {
	Date      time.Time `json: "Date"`
	Selection string    `json: "Selection"`
}

type User struct {
	Name    string   `json: "Name"`
	Choices []Choice `json: "Choices"`
}

type Pickem struct {
	Name    string    `json: "Name"`
	Option1 string    `json: "Option1"`
	Option2 string    `json: "Option2"`
	Date    time.Time `json: "Date"`
	Winner  string    `json: "Winner"`
}

type Group struct {
	Name    string   `json: "Name"`
	ID      string   `json: "ID"`
	Users   []User   `json: "Users"`
	Pickems []Pickem `json: "Pickems"`
}

func upsertGroup(db *mongo.Database, coll *mongo.Collection, group Group) {
	opts := options.Update().SetUpsert(true)
	upsertdata := bson.M{"$set": group}
	filter := bson.D{{"ID", group.ID}}

	result, err := coll.UpdateOne(
		context.Background(),
		filter,
		upsertdata,
		opts,
	)

	if err != nil {
		fmt.Println("Error Upserting row: " + err.Error())
		return
	}

	if result.MatchedCount != 0 {
		fmt.Println("Updated existing document")
		return
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("Inserted a new document with ID %v\n", result.UpsertedID)
	}
}

func getAllGroups(db *mongo.Database, coll *mongo.Collection) ([]Group, error) {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{},
	)

	if err != nil {
		return nil, errors.New("Error Finding collection: " + err.Error())
	}

	var groups []Group

	for cursor.Next(context.TODO()) {
		elem := &Group{}
		if err := cursor.Decode(elem); err != nil {
			log.Fatal(err)
		}
		// ideally, you would do something with elem....
		// but for now just print it to the console
		// fmt.Println(elem)
		groups = append(groups, *elem)
	}

	return groups, nil
}

func getGroupWithID(db *mongo.Database, coll *mongo.Collection, ID string) (Group, error) {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{{"ID", ID}},
	)

	if err != nil {
		return Group{}, errors.New("Error Finding collection: " + err.Error())
	}

	var group Group

	for cursor.Next(context.TODO()) {
		elem := &Group{}
		if err := cursor.Decode(elem); err != nil {
			log.Fatal(err)
		}
		// ideally, you would do something with elem....
		// but for now just print it to the console
		// fmt.Println(elem)
		group = *elem
	}

	return group, nil
}

// func createGroup(name string)

func readAll(db *mongo.Database, coll *mongo.Collection) {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{},
	)

	if err != nil {
		fmt.Println("Error Finding collection: " + err.Error())
	}

	for cursor.Next(context.TODO()) {
		elem := &bson.D{}
		if err := cursor.Decode(elem); err != nil {
			log.Fatal(err)
		}
		// ideally, you would do something with elem....
		// but for now just print it to the console
		fmt.Println(elem)
	}

}

func deleteAll(db *mongo.Database, coll *mongo.Collection) {

	err := coll.Drop(context.TODO())

	if err != nil {
		fmt.Println("Error deleting all rows: " + err.Error())
	}

	fmt.Println("Successfully deleted all")
}

func genID() (string, error) {
	ID := make([]byte, 12)

	_, err := rand.Read(ID)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(ID), nil
}

func main() {
	connectionString := os.Getenv("PickemConnectionString")
	if connectionString == "" {
		log.Fatal("Please set PickemConnectionString")
	}

	// Open Connection
	clientOpts := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB Luigi cluster!")

	db := client.Database("pickem")
	coll := db.Collection("groups")

	const shortForm = "2006-Jan-02"
	yesterDay, _ := time.Parse(shortForm, "2019-Nov-18")

	exampleGroup := Group{
		Name: "P and P",
		ID:   "2",
		Users: []User{
			User{
				Name: "Pratique",
				Choices: []Choice{
					Choice{
						Date:      time.Now(),
						Selection: "Pizza",
					},
					Choice{
						Date:      yesterDay,
						Selection: "Hotdog",
					},
				},
			},
			User{
				Name: "Pagna",
				Choices: []Choice{
					Choice{
						Date:      time.Now(),
						Selection: "Burger",
					},
					Choice{
						Date:      yesterDay,
						Selection: "Corndog",
					},
				},
			},
		},
		Pickems: []Pickem{
			Pickem{
				Name:    "Burger vs Pizza",
				Option1: "Burger",
				Option2: "Pizza",
				Date:    time.Now(),
				Winner:  "Pizza",
			},
			Pickem{
				Name:    "Hotdog vs Corndog",
				Option1: "Hotdog",
				Option2: "Corndog",
				Date:    time.Now(),
				Winner:  "Hotdog",
			},
		},
	}

	upsertGroup(db, coll, exampleGroup)
	// deleteAll(db, coll)
	// readAll(db, coll)

	group, _ := getGroupWithID(db, coll, "1")

	fmt.Println(group.Name)

	// groups, err := getAllGroups(db, coll)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	//
	// for _, group := range groups {
	// 	fmt.Println(group.Name)
	// }

}
