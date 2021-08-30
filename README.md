# Hash identifier

This is fairly basic and is used to try and identify if a string has any specific hash applied to it. E.g. is this string likely a `MD5` hash?

Example of useage:
```go
hash := "$mongodb$0$sa$75692b1d11c072c6c79332e248c4f699"
listOfMatchingHashes, _ := identify(hash)
```

This will return `[]Hash` which looks like:
```go
type Hash struct {
    Name string
    John string
    Hashcat int32
    Extended bool
}
```
