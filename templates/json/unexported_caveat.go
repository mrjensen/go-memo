package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Secret struct {
	Public  string `json:"public"`
	private string `json:"private"`
}

func main() {
	secret := Secret{
		Public:  "public",
		private: "private",
	}
	data, err := json.MarshalIndent(secret, "", "\t")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%s", data)
}
