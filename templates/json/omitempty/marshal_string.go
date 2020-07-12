package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName,omitempty"`
	Alias     *string `json:"alias"`
	Nickname  *string `json:"nickname,omitempty"`
}

func main() {
	user := User{
		FirstName: "",
		LastName:  "",
		Alias:     nil,
		Nickname:  nil,
	}

	data, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%s", data)
}
