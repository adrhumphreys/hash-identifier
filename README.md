# Hash identifier

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
