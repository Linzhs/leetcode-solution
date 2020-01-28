package pkg

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	nums := []int{1, 3, 1, 4, 2, 4, 2, 3, 2, 4, 7, 6, 6, 7, 5, 5, 7, 7}
	fmt.Printf("%+v", insertionSort(nums))
}
