package models

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
	UID     string   `json: "UID"`
	Choices []Choice `json: "Choices"`
	IsAdmin bool     `json: "IsAdmin"`
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

func UpsertGroup(db *mongo.Database, coll *mongo.Collection, group Group) error {
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
		return errors.New("Error Upserting row: " + err.Error())
	}

	if result.MatchedCount != 0 {
		fmt.Println("Updated existing document")
		return nil
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("Inserted a new document with ID %v\n", result.UpsertedID)
		return nil
	}
	return nil
}

func GetAllGroups(db *mongo.Database, coll *mongo.Collection) ([]Group, error) {
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

func GetGroupWithID(db *mongo.Database, coll *mongo.Collection, ID string) (Group, error) {
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

func ReadAll(db *mongo.Database, coll *mongo.Collection) {
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

func DeleteAll(db *mongo.Database, coll *mongo.Collection) {

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

func Connect() (*mongo.Client, error) {
	connectionString := os.Getenv("PickemConnectionString")
	if connectionString == "" {
		log.Fatal("Please set PickemConnectionString")
	}

	// Open Connection
	clientOpts := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		return &mongo.Client{}, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return &mongo.Client{}, err
	}

	return client, err
}

func CreateGroup(groupName string, adminUID string) (Group, error) {
	if groupName == "" {
		return Group{}, errors.New("Empty group name")
	}

	if adminUID == "" {
		return Group{}, errors.New("Empty admin name")
	}

	id, err := genID()
	if err != nil {
		return Group{}, errors.New("Error generating id: " + err.Error())
	}
	group := Group{
		Name: groupName,
		ID:   id,
		Users: []User{
			User{
				UID:     adminUID,
				Choices: []Choice{},
				IsAdmin: true,
			},
		},
		Pickems: []Pickem{},
	}

	return group, nil
}

func CreateUser(UID string) User {
	return User{
		UID:     UID,
		Choices: []Choice{},
		IsAdmin: false,
	}

}

func AddUserToGroup(group Group, user User) (Group, error) {

	group.Users = append(group.Users, user)

	return group, nil
}
