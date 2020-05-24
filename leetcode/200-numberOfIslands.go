package leetcode

type unionFindSet struct {
	roots []int
	count int
}

func newUnionFindSetByGrid(grid [][]byte) *unionFindSet {
	m, n := len(grid), len(grid[0])

	ufSet := &unionFindSet{
		roots: make([]int, m*n),
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			root := i*n + j
			if grid[i][j] == '1' {
				ufSet.roots[root] = root
				ufSet.count++
				continue
			}
			ufSet.roots[root] = -1
		}
	}

	return ufSet
}

func (s *unionFindSet) findRoot(i int) int {
	root := i

	// find root
	for root != s.roots[root] {
		root = s.roots[root]
	}

	// path compression
	for i != s.roots[root] {
		i, s.roots[i] = s.roots[i], root
	}

	return root
}

func (s *unionFindSet) connected(p, q int) bool {
	return s.findRoot(p) == s.findRoot(q)
}

func (s *unionFindSet) union(p, q int) {
	pRoot := s.findRoot(p)
	qRoot := s.findRoot(q)
	if pRoot == qRoot {
		return
	}

	s.roots[pRoot] = qRoot
	s.count--
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var (
		m, n       = len(grid), len(grid[0])
		directions = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
		ufSet      = newUnionFindSetByGrid(grid)
	)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}

			for _, d := range directions {
				nr, nc := i+d[0], j+d[1]
				if nc >= 0 && nr >= 0 && nr < m && nc < n && grid[nr][nc] == '1' {
					ufSet.union(i*n+j, nr*n+nc)
				}
			}
		}
	}

	return ufSet.count
}
