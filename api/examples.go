package examples

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func writeOne(db *mongo.Database, coll *mongo.Collection) {

	result, err := coll.InsertOne(
		context.Background(),
		bson.D{
			{"item", "canvas"},
			{"qty", 100},
			{"tags", bson.A{"cotton"}},
			{"size", bson.D{
				{"h", 28},
				{"w", 35.5},
				{"uom", "cm"},
			}},
		})

	if err != nil {
		fmt.Println("Error Inserting row: " + err.Error())
	}

	fmt.Printf("Inserted document with ID %v\n", result.InsertedID)
}

func read(db *mongo.Database, coll *mongo.Collection) {
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

func writeMany(db *mongo.Database, coll *mongo.Collection) {
	docs := []interface{}{
		bson.D{
			{"item", "journal"},
			{"qty", 25},
			{"size", bson.D{
				{"h", 14},
				{"w", 21},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "notebook"},
			{"qty", 50},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "paper"},
			{"qty", 100},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "planner"},
			{"qty", 75},
			{"size", bson.D{
				{"h", 22.85},
				{"w", 30},
				{"uom", "cm"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "postcard"},
			{"qty", 45},
			{"size", bson.D{
				{"h", 10},
				{"w", 15.25},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
	}

	result, err := coll.InsertMany(context.Background(), docs)

	if err != nil {
		fmt.Println("Error Inserting row: " + err.Error())
	}

	for insertID := range result.InsertedIDs {
		fmt.Printf("Inserted document with ID %v\n", insertID)
	}

}

func query(db *mongo.Database, coll *mongo.Collection) {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{{"status", "D"}},
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

func queryExact(db *mongo.Database, coll *mongo.Collection) {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{
			{"size", bson.D{
				{"h", 14},
				{"w", 21},
				{"uom", "cm"},
			}},
		},
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

func queryDot(db *mongo.Database, coll *mongo.Collection) {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{{"size.uom", "in"}},
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

func queryLessThan(db *mongo.Database, coll *mongo.Collection) {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{
			{"size.h", bson.D{
				{"$lt", 15},
			}},
		},
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

func queryAnd(db *mongo.Database, coll *mongo.Collection) {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{
			{"status", "A"},
			{"qty", bson.D{{"$lt", 30}}},
		},
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

func queryOr(db *mongo.Database, coll *mongo.Collection) {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{
			{"$or",
				bson.A{
					bson.D{{"status", "A"}},
					bson.D{{"qty", bson.D{{"$lt", 30}}}},
				}},
		},
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

func queryAndOr(db *mongo.Database, coll *mongo.Collection) {
	cursor, err := coll.Find(
		context.Background(),
		bson.D{
			{"status", "A"},
			{"$or", bson.A{
				bson.D{{"qty", bson.D{{"$lt", 30}}}},
				bson.D{{"item", primitive.Regex{Pattern: "^p", Options: ""}}},
			}},
		},
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

func updateOne(db *mongo.Database, coll *mongo.Collection) {

	result, err := coll.UpdateOne(
		context.Background(),
		bson.D{
			{"item", "paper"},
		},
		bson.D{
			{"$set", bson.D{
				{"size.uom", "cm"},
				{"status", "P"},
			}},
			{"$currentDate", bson.D{
				{"lastModified", true},
			}},
		},
	)

	if err != nil {
		fmt.Println("Error updating row: " + err.Error())
	}

	fmt.Printf("Updated document with ID %v\n", result.UpsertedID)
}

func updateMany(db *mongo.Database, coll *mongo.Collection) {
	_, err := coll.UpdateMany(
		context.Background(),
		bson.D{
			{"qty", bson.D{
				{"$lt", 50},
			}},
		},
		bson.D{
			{"$set", bson.D{
				{"size.uom", "cm"},
				{"status", "P"},
			}},
			{"$currentDate", bson.D{
				{"lastModified", true},
			}},
		},
	)

	if err != nil {
		fmt.Println("Error updating row: " + err.Error())
	}

}

func deleteOne(db *mongo.Database, coll *mongo.Collection) {

	_, err := coll.DeleteOne(
		context.Background(),
		bson.D{
			{"status", "D"},
		},
	)

	if err != nil {
		fmt.Println("Error deleting row: " + err.Error())
	}

	fmt.Println("Successfully deleted")
}

func deleteMany(db *mongo.Database, coll *mongo.Collection) {

	_, err := coll.DeleteMany(
		context.Background(),
		bson.D{
			{"item", "*"},
		},
	)

	if err != nil {
		fmt.Println("Error deleting rows: " + err.Error())
	}

	fmt.Println("Successfully deleted many")
}

func deleteAll(db *mongo.Database, coll *mongo.Collection) {

	err := coll.Drop(context.TODO())

	if err != nil {
		fmt.Println("Error deleting all rows: " + err.Error())
	}

	fmt.Println("Successfully deleted all")
}

func main() {
	// Open Connection
	clientOpts := options.Client().ApplyURI("mongodb+srv://dylan:Yu92RzKzy7GYEqE1@luigi-aoo0c.mongodb.net/test?retryWrites=true")
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatal(err)
	}

	// End Open Connection Code

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database("test")
	coll := db.Collection("inventory")

	// write(db, coll)
	// read(db, coll)
	// writeMany(db, coll)
	// query(db, coll)
	// queryDot(db, coll)
	// queryExact(db, coll)
	// queryLessThan(db, coll)
	// queryAnd(db, coll)
	// queryOr(db, coll)
	// queryAndOr(db, coll)
	// updateOne(db, coll)
	// updateMany(db, coll)
	// deleteOne(db, coll)
	// deleteMany(db, coll)
	deleteAll(db, coll)
	read(db, coll)

}
