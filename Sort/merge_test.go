package Sort

import (
	"fmt"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func Test_Merge_Sort(t *testing.T) {
	a := []int{6, 3, 2, 4, 8, 9, 1, 5, 7}
	MergeSort(a)
	fmt.Println(a)
	//assert.
}
