# Unexported caveat

Unexported fields in a struct are not marshalled to JSON (`private` field in this case).

JSON marshalling the following struct:
```go
secret := Secret{
    Public:  "public",
    private: "private",
}
```
outputs:

```
{
	"public": "public"
}
```

See file: [/../../blob/master/templates/json/unexported_caveat.go](/../../blob/master/templates/json/unexported_caveat.go)