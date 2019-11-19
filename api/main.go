package main

import (
	"fmt"
	"log"
	"net/http"

	"models/models"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func handleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db := client.Database("pickem")
		coll := db.Collection("groups")

		models.DeleteAll(db, coll)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Invalid Method\n")
	}
}

func handleRead(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db := client.Database("pickem")
		coll := db.Collection("groups")

		models.ReadAll(db, coll)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Invalid Method\n")
	}
}

func handleAddGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		groupName := r.Form["groupName"]
		if len(groupName) == 0 {
			return
		}
		adminUID := r.Form["adminUID"]
		if len(adminUID) == 0 {
			return
		}

		group, err := models.CreateGroup(groupName[0], adminUID[0])
		if err != nil {
			fmt.Println(err.Error())
		}

		db := client.Database("pickem")
		coll := db.Collection("groups")
		err = models.UpsertGroup(db, coll, group)
		if err != nil {
			fmt.Println(err.Error())
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Invalid Method\n")
	}
}

func init() {
	var err error
	client, err = models.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to MongoDB Luigi cluster!")
}

func main() {

	http.HandleFunc("/delete", handleDelete)
	http.HandleFunc("/addGroup", handleAddGroup)
	http.HandleFunc("/read", handleRead)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
