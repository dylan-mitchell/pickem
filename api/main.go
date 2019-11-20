package main

import (
	"crypto/rand"
	"encoding/hex"
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

func genID() (string, error) {
	ID := make([]byte, 12)

	_, err := rand.Read(ID)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(ID), nil
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

		id, err := genID()
		if err != nil {
			fmt.Println("Error generating ID")
			return
		}

		db := client.Database("pickem")
		coll := db.Collection("groups")

		exists := models.CheckIfIDExists(db, coll, id)
		for exists {
			fmt.Println("ID already exists")
			fmt.Println("Generating new one")

			id, err = genID()
			if err != nil {
				fmt.Println("Error generating ID")
				return
			}
			exists = models.CheckIfIDExists(db, coll, id)
		}

		group, err := models.CreateGroup(groupName[0], adminUID[0], id)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = models.UpsertGroup(db, coll, group)
		if err != nil {
			fmt.Println(err.Error())
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Invalid Method\n")
	}
}

func handleGetGroupsForUID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		UID := r.Form["UID"]
		if len(UID) == 0 {
			return
		}

		db := client.Database("pickem")
		coll := db.Collection("groups")

		groups, err := models.GetGroupsForUID(db, coll, UID[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if len(groups) == 0 {
			fmt.Println("No groups found for " + UID[0])
		} else {
			for _, group := range groups {
				fmt.Println(group.ID)
			}
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Invalid Method\n")
	}
}

func handleAddUserToGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		ID := r.Form["ID"]
		if len(ID) == 0 {
			return
		}
		UID := r.Form["UID"]
		if len(UID) == 0 {
			return
		}

		db := client.Database("pickem")
		coll := db.Collection("groups")

		err := models.AddUserToGroupWithID(db, coll, ID[0], UID[0])
		if err != nil {
			fmt.Println(err.Error())
			return
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
	http.HandleFunc("/getGroupsForUID", handleGetGroupsForUID)
	http.HandleFunc("/addUserToGroup", handleAddUserToGroup)
	http.HandleFunc("/read", handleRead)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
