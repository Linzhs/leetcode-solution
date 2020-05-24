package leetcode

func findOrder(numCourses int, prerequisites [][]int) (result []int) {
	// 构建入度数组和邻接表
	var (
		inDegrees     = make([]int, numCourses)
		adjacencyList = make([][]int, numCourses)
	)
	for _, list := range prerequisites {
		inDegrees[list[0]]++
		adjacencyList[list[1]] = append(adjacencyList[list[1]], list[0])
	}

	// 将入度为0的节点入队
	var queue []int
	for i, v := range inDegrees {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	// 拓扑排序
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]

		numCourses--
		result = append(result, elem)
		for _, v := range adjacencyList[elem] {
			inDegrees[v]--
			if inDegrees[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if numCourses == 0 {
		return
	}

	return nil
}
