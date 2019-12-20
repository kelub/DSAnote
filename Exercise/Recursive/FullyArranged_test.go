package Recursive

import (
	"testing"
)

func Test_Fully(t *testing.T) {
	a := []int{1, 2, 3, 4}
	lena := len(a)
	Fully(a, lena, lena)
}
