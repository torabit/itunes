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

### API Examples
There is an example of the API in the examples directory
