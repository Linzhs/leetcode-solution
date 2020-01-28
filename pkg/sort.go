package pkg

func bubbleSort(nums []int) []int {
	return nil
}

func insertionSort(nums []int) []int {
	sortedSlice := make([]int, 0, len(nums))
	sortedSlice = append(sortedSlice, nums...)

	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if sortedSlice[j-1] > sortedSlice[j] {
				sortedSlice[j-1], sortedSlice[j] = sortedSlice[j], sortedSlice[j-1]
				continue
			}
			break
		}
	}

	return sortedSlice
}
