package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	struct1 := Struct1{}
	err := json.Unmarshal([]byte(data), &struct1)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%+v", struct1)
}

type Struct1 struct {
	String1 string  `json:"string1"`
	Number1 int     `json:"number1"`
	Object1 Object1 `json:"object1"`
}

type Object1 struct {
	Nested string `json:"nested"`
}

var data = `
{
    "string1": "string1",
    "number1": 1,
    "object1": {
        "nested": "nested1"
    },
    "string2": "string2",
    "number2": 2,
    "object2": {
        "nested": "nested2"
    }
}
`
