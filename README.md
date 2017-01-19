## [Selection](https://en.wikipedia.org/wiki/Selection_algorithm)

[![GoDoc](https://godoc.org/github.com/chasestarr/selection?status.svg)](https://godoc.org/github.com/chasestarr/selection)

```go
import (
  "fmt"

  "github.com/chasestarr/selection"
)

func main() {
  input := []int{3, 8, 2, 5, 1, 4, 7, 6}
  // returns the ith smallest array element
  s := selection.Select(input, 4)
  fmt.Println(s) // 4
}
```