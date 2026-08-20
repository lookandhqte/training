package coderun

import "fmt"

func maxCostPath() {
	var n, m int
	fmt.Scan(&n, &m)
	grid := make([][]int, n)

	for i := range grid {
		grid[i] = make([]int, m)
		for j := range grid[i] {
			fmt.Scan(&grid[i][j])
		}
	}

	dm := make([][]int, n)
	for i := range dm {
		dm[i] = make([]int, m)
	}

	dm[0][0] = grid[0][0]

	for i := 1; i < n; i++ {
		dm[i][0] = dm[i-1][0] + grid[i][0]
	}

	for j := 1; j < m; j++ {
		dm[0][j] = dm[0][j-1] + grid[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if dm[i-1][j] > dm[i][j-1] {
				dm[i][j] = dm[i-1][j] + grid[i][j]
			} else {
				dm[i][j] = dm[i][j-1] + grid[i][j]
			}
		}
	}

	path := make([]byte, 0, n+m-2)
	r, c := n-1, m-1
	for r > 0 || c > 0 {
		if r == 0 {
			path = append(path, 'R')
			c--
		} else if c == 0 {
			path = append(path, 'D')
			r--
		} else if dm[r-1][c] > dm[r][c-1] {
			path = append(path, 'D')
			r--
		} else {
			path = append(path, 'R')
			c--
		}
	}

	if len(path) > 0 {
		for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}

		result := make([]byte, 0, len(path)*2-1)
		result = append(result, path[0])
		l := len(path)
		for i := 1; i < l; i++ {
			result = append(result, ' ', path[i])
		}

		fmt.Println(dm[n-1][m-1])
		fmt.Println(string(result))
	} else {
		fmt.Println(dm[0][0])
		fmt.Println()
	}

}
