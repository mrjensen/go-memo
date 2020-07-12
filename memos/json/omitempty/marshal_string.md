# String
Struct with zero values for strings and string pointers
```go
type User struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName,omitempty"`
	Alias     *string `json:"alias"`
	Nickname  *string `json:"nickname,omitempty"`
}
```
outputs:
```
{
	"firstName": "",
	"alias": null
}
```

See file: [/../../edit/master/templates/json/omitempty/marshal_string.go](/../../edit/master/templates/json/omitempty/marshal_string.go)