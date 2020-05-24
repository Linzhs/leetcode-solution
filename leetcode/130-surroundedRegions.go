package leetcode

type ufsRegion []int

func newUfsRegion(board [][]byte) ufsRegion {
	m, n := len(board), len(board[0])

	ufs := make(ufsRegion, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			root := i*n + j
			if board[i][j] == 'X' {
				ufs[root] = 0
				continue
			}
			ufs[root] = root
		}
	}

	return ufs
}

func (s ufsRegion) findRoot(i int) int {
	root := i

	// find root
	for root != s[root] {
		root = s[root]
	}

	// path compression
	for i != s[root] {
		i, s[i] = s[i], root
	}

	return root
}

func (s ufsRegion) connected(p, q int) bool {
	return s.findRoot(p) == s.findRoot(q)
}

func (s ufsRegion) unionO(p, q int) {
	pRoot := s.findRoot(p)
	qRoot := s.findRoot(q)
	if pRoot == qRoot {
		return
	}

	s[pRoot] = qRoot
}

// union & find
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	var (
		m, n       = len(board), len(board[0])
		directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		ufs        = newUfsRegion(board)
	)

	fnMark := func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n && board[i][j] == 'X' {
			return
		}

		for _, d := range directions {
			nr, nc := i+d[0], j+d[1]
			if nr >= 0 && nc >= 0 && nr < m && nc < n && board[nr][nc] == 'O' {
				ufs.unionO(i*n+j, nr*n+nc)
			}
		}
	}

	for i := 0; i < m; i++ {
		fnMark(i, 0)
		fnMark(i, n-1)
	}

	for j := 0; j < n; j++ {
		fnMark(0, j)
		fnMark(m-1, j)
	}

	for i, v := range ufs {
		if v == 0 {
			board[i/n][i%n] = 'X'
		}
	}
}

func solveDfs(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	if len(board) < 2 || len(board[0]) < 2 {
		return
	}

	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			boundaryDfs(board, i, 0)
		}
		if board[i][n-1] == 'O' {
			boundaryDfs(board, i, n-1)
		}
	}

	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			boundaryDfs(board, 0, j)
		}
		if board[m-1][j] == 'O' {
			boundaryDfs(board, m-1, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}
}

func boundaryDfs(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 {
		return
	}

	if board[i][j] == 'O' {
		board[i][j] = '*'
	}
	if i > 1 && board[i-1][j] == 'O' {
		boundaryDfs(board, i-1, j)
	}
	if i < len(board)-2 && board[i+1][j] == 'O' {
		boundaryDfs(board, i+1, j)
	}
	if j > 1 && board[i][j-1] == 'O' {
		boundaryDfs(board, i, j-1)
	}
	if j < len(board[0])-1 && board[i][j+1] == 'O' {
		boundaryDfs(board, i, j+1)
	}
}
