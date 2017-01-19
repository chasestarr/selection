package selection

import (
	"fmt"
	"math/rand"
	"time"
)

// Select returns the ith smallest int from an int slice
func Select(a []int, i int) int {
	if i < 0 || i >= len(a) {
		return -1
	}

	p := partition(a, 0, len(a))
	for p+1 != i {
		fmt.Println(p)
		if p+1 < i {
			p = partition(a, p+1, len(a))
		} else {
			p = partition(a, 0, p)
		}
	}

	return a[p]
}

func partition(a []int, left, right int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	p := r1.Intn(right-left) + left
	pivot := a[p]
	swap(a, left, p)
	i := left + 1
	for j := left + 1; j < right; j++ {
		if a[j] < pivot {
			swap(a, i, j)
			i++
		}
	}
	swap(a, i-1, left)
	return i - 1
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
