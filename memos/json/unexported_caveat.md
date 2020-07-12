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

See file: [/../../edit/master/templates/json/unexported_caveat.go](/../../edit/master/templates/json/unexported_caveat.go)