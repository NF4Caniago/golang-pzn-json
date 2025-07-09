package golangpznjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"Kurniawan","LastName":"Khanendy","Age":30,"Married":true,"Hobbies":["Game","Kuliner","Traveling"],"Addresses":[{"Street":"Tebet","Country":"Indonesia","PostalCode":"123345"},{"Street":"Rawamangun","Country":"Zimbabwe","PostalCode":"123333"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Addresses)
}
