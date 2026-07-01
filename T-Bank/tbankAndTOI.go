package tbank

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isTOI(c byte) bool {
	return c == 'T' || c == 'O' || c == 'I'
}

func tbankAndTOI() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][i] = 0
	}
	for i := 0; i < n; i++ {
		if isTOI(s[i]) {
			dp[i][i+1] = 0
		} else {
			dp[i][i+1] = 1
		}
	}
	for length := 2; length <= n; length++ {
		for l := 0; l+length <= n; l++ {
			r := l + length
			best := 1 + min(dp[l+1][r], dp[l][r-1])
			cost := 0
			if !isTOI(s[l]) {
				cost++
			}
			if !isTOI(s[r-1]) {
				cost++
			}
			if isTOI(s[l]) && isTOI(s[r-1]) && s[l] != s[r-1] {
				cost++
			}
			best = min(best, cost+dp[l+1][r-1])
			dp[l][r] = best
		}
	}

	fmt.Println(dp[0][n])
}
