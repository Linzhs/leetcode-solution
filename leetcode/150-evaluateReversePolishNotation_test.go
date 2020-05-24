package leetcode

import (
	"fmt"
	"testing"
)

func TestEvalRpn(t *testing.T) {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
