package leetcode

// 拓扑排序 入度表 广度优先搜索
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 构造入度表和邻接表
	var (
		adjacencyList = make([][]int, numCourses)
		inDegrees     = make([]int, numCourses)
	)
	for _, list := range prerequisites {
		inDegrees[list[0]]++
		adjacencyList[list[1]] = append(adjacencyList[list[1]], list[0])
	}

	// 找出入度为0的节点并入队
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
		for _, cur := range adjacencyList[elem] {
			inDegrees[cur]--
			if inDegrees[cur] == 0 {
				queue = append(queue, cur)
			}
		}
	}

	return numCourses == 0
}
