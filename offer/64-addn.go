package offer

func sumNums(n int) (result int) {
	return sumNumsHelper(&result, n)
}

func sumNumsHelper(result *int, n int) int {
	_ = n > 1 && sumNumsHelper(result, n-1) > 0
	*result += n
	return *result
}
