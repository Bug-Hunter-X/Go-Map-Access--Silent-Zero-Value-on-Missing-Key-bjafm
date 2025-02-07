```go
func main() {
    var m map[string]int
    fmt.Println(m["a"]) // this will not panic, but returns 0.  This can be unexpected.
}
```