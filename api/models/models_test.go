package models

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	_, err := Connect()
	if err != nil {
		t.Fatalf("Failed connecting to Database")
	}
	fmt.Println("Connected to MongoDB Luigi cluster!")
}

type input struct {
	groupName string
	adminUID  string
}

func TestCreateGroup(t *testing.T) {
	tests := []struct {
		name      string
		input     input
		shouldErr bool
	}{
		{
			name:      "Test valid Create Group",
			input:     input{groupName: "test", adminUID: "Test admin"},
			shouldErr: false,
		},
		{
			name:      "Test empty groupname Create Group",
			input:     input{groupName: "", adminUID: "Test admin"},
			shouldErr: true,
		},
		{
			name:      "Test empty admin Create Group",
			input:     input{groupName: "test", adminUID: ""},
			shouldErr: true,
		},
	}

	for _, test := range tests {
		group, err := CreateGroup(test.input.groupName, test.input.adminUID)
		if err != nil && !test.shouldErr {
			t.Fatalf("Test errored when it should not have")
		}
		if err == nil && test.shouldErr {
			t.Fatalf("Test did not error when it should have")
		}
		if err == nil {
			if group.Users[0].IsAdmin != true {
				t.Fatalf("User who creates group should be ")
			}
		}
		fmt.Println(test.name + " passed")
	}
}
