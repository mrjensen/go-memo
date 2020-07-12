# Unmarshal
Unmarshalling JSON maps the known fields from a struct and ignores the rest.

For the struct:
```go
type Struct1 struct {
	String1 string  `json:"string1"`
	Number1 int     `json:"number1"`
	Object1 Object1 `json:"object1"`
}

type Object1 struct {
	Nested string `json:"nested"`
}
```
and the input:
```json
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
```
the printed struct is (format: `%+v`):
```
{String1:string1 Number1:1 Object1:{Nested:nested1}}
```

See file: [templates/json/omitempty/unmarshal.go](templates/json/omitempty/unmarshal.go)