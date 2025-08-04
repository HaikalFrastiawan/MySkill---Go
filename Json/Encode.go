package main

import (
	"encoding/json"
	"fmt"
)

type User struct {  	
	FullName string `json:"Name"`
	Age      int
}

func main() {


	var object =  []User{{"jhon wick", 27}, {"ethan", 10}}
	var jsonData, err = json.Marshal(object) // untuk nampilin json
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)

}