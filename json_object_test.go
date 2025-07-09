package golangpznjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khanendy",
		Age:        30,
		Married:    true,
		Hobbies: []string{
			"Game",
			"Kuliner",
			"Traveling",
		},
		Addresses: []Address{
			{
				Street:     "Tebet",
				Country:    "Indonesia",
				PostalCode: "123345",
			},
			{
				Street:     "Rawamangun",
				Country:    "Zimbabwe",
				PostalCode: "123333",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
