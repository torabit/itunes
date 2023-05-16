# iTunes

This is a Go wrapper for working with [iTuness Search API](https://developer.apple.com/library/archive/documentation/AudioVideo/Conceptual/iTuneSearchAPI/index.html#//apple_ref/doc/uid/TP40017632-CH3-SW1).

## Installation
```
go get github.com/torabit/itunes
```
## Package Usage

### import
```go
import . "github.com/torabit/itunes"
```

### Search Example
```go
ctx := context.Background()
client := New()

results, err := client.Search(
  ctx,
  Term("Jamiroquai"),
  Media(MediaMusic),
  Limit(5)
)
```
results has `ResultCount` and `Result[]`

### Lookup Example
```go
ctx := context.Background()
client := New()

results, err := client.Lookup(
  ctx,
  ID(1611115294)
)
```
Lookup method requires LookupOption as an argument  
You can also pass requestOption as an argument if necessary
### API Examples
There is an example of the API in the [examples](https://github.com/torabit/itunes/tree/master/examples) directory
