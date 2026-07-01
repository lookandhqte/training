package tbank

import "fmt"

const MOD = 1000000007

func tbankAndParty() {
	var n, m int
	fmt.Scan(&n, &m)
	res := (m * (m - 1)) % MOD
	base := ((m*m-3*m+3)%MOD + MOD) % MOD
	pow := n - 1
	for pow > 0 {
		if pow&1 == 1 {
			res = (res * base) % MOD
		}
		base = (base * base) % MOD
		pow >>= 1
	}
	fmt.Println(res)
}
