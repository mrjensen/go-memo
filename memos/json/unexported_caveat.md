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

See file: [templates/json/unexported_caveat.go](templates/json/unexported_caveat.go)