package models

import (
	"context"
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
	Date      time.Time `json: "date"`
	Selection string    `json: "selection"`
	User      string    `json: "user"`
}

type Pickem struct {
	Name    string    `json: "name"`
	Option1 string    `json: "option1"`
	Option2 string    `json: "option2"`
	Date    time.Time `json: "date"`
	Winner  string    `json: "winner"`
	Choices []Choice  `json: "choices"`
}

type Group struct {
	Name    string   `json: "name"`
	ID      string   `json: "id"`
	Users   []string `json: "users"`
	Pickems []Pickem `json: "pickems"`
	Admin   string   `json: "admin"`
}

// func (g Group) String() string {
// 	var groupString string
// 	groupString = groupString + "*****************\n"
// 	groupString = groupString + "ID: " + g.ID + "\n"
// 	groupString = groupString + "Name: " + g.Name + "\n"
// 	for i, pickem := range g.Pickems {
// 		groupString = groupString + fmt.Sprintf("Pickem%d:\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n", i, pickem.Name, pickem.Option1, pickem.Option2, pickem.Winner, pickem.Date)
// 	}
// 	for i, user := range g.Users {
// 		var adminStr string
// 		if user.IsAdmin {
// 			adminStr = "isAdmin"
// 		} else {
// 			adminStr = "isNotAdmin"
// 		}
// 		groupString = groupString + fmt.Sprintf("User%d:\n\t%s\n\t%s\n", i, user.UID, adminStr)
// 		for i, choice := range user.Choices {
// 			groupString = groupString + fmt.Sprintf("Choice%d:\n\t%s\n\t%s\n", i, choice.Selection, choice.Date)
// 		}
// 	}
// 	groupString = groupString + "*****************\n"
//
// 	return groupString
// }

func UpsertGroup(db *mongo.Database, coll *mongo.Collection, group Group) error {
	opts := options.Update().SetUpsert(true)
	upsertdata := bson.M{"$set": group}
	filter := bson.D{{"id", group.ID}}

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
		bson.D{{"id", ID}},
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

		group = *elem
	}

	return group, nil
}

func AddUserToGroupWithID(db *mongo.Database, coll *mongo.Collection, ID string, UID string) error {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{{"id", ID}},
	)

	if err != nil {
		return errors.New("Error Finding collection: " + err.Error())
	}

	var group Group

	for cursor.Next(context.TODO()) {
		elem := &Group{}
		if err := cursor.Decode(elem); err != nil {
			return err
		}

		group = *elem
	}

	group.Users = append(group.Users, UID)
	UpsertGroup(db, coll, group)
	return nil
}

func GetGroupsForUID(db *mongo.Database, coll *mongo.Collection, UID string) ([]Group, error) {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{
			{"users", UID},
		},
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

		groups = append(groups, *elem)
	}

	return groups, nil
}

func ReadAll(db *mongo.Database, coll *mongo.Collection) {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{},
	)

	if err != nil {
		fmt.Println("Error Finding collection: " + err.Error())
	}

	for cursor.Next(context.TODO()) {
		elem := &Group{}
		if err := cursor.Decode(elem); err != nil {
			log.Fatal(err)
		}

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

func CheckIfIDExists(db *mongo.Database, coll *mongo.Collection, id string) bool {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{{"id", id}},
	)

	if err != nil {
		fmt.Println("Error Finding collection: " + err.Error())
	}

	for cursor.Next(context.TODO()) {
		elem := &bson.D{}
		if err := cursor.Decode(elem); err != nil {
			log.Fatal(err)
		}
		return true
	}
	return false

}

func CreateGroup(groupName string, adminUID string, id string) (Group, error) {
	if groupName == "" {
		return Group{}, errors.New("Empty group name")
	}

	if adminUID == "" {
		return Group{}, errors.New("Empty admin name")
	}

	group := Group{
		Name:    groupName,
		ID:      id,
		Users:   []string{adminUID},
		Admin:   adminUID,
		Pickems: []Pickem{},
	}

	return group, nil
}

func AddUserToGroup(group Group, user string) (Group, error) {

	group.Users = append(group.Users, user)

	return group, nil
}
