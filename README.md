# complexconv

ParseComplex returns the complex value represented by the string.

```go
import complexconv

func main() {
	fmt.Println(complexconv.ParseComplex("3.14159 + 2.71828i", 128))
}
```

FormatComplex returns the string representation of complex.

```go
import complexconv

func main() {
	fmt.Println(complexconv.FormatComplex(3.14159 + 2.71828i))
}
```
