package api

import (
	"fmt"
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	fmt.Println(fractionToDecimal(1, 3))
	fmt.Println(fractionToDecimal(4, 333))
}
