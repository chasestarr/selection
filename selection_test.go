package selection

import "testing"

func TestSimple(t *testing.T) {
	input := []int{1, 2, 4, 3, 5}
	i := 3
	correct := 3
	if output := Select(input, i); output != correct {
		t.Fatal("incorrect ith element... expected:", correct, "received:", output)
	}
}

func TestLonger(t *testing.T) {
	input := []int{3, 8, 2, 5, 1, 4, 7, 6}
	i := 4
	correct := 4
	if output := Select(input, i); output != correct {
		t.Fatal("incorrect ith element... expected:", correct, "received:", output)
	}
}
