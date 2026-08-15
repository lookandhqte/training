package leetcode

// _* = 0 or more prev characters
// . = 1 any character
// .* = repeat any character zero or more times

func isMatch(s string, p string) bool {
	if s == p || p == ".*" {
		return true
	}
	pid := 0
	i := 0
	for i < len(s) {
		if pid == len(p) {
			if i < len(s) {
				return false
			}
		}
		if s[i] == p[pid] || p[pid] == '.' {
			i++
			pid++
			continue
		} else if p[pid] == '*' {
			elem := p[pid-1] //a
			for i < len(s) {
				if s[i] == elem {
					i++
				} else {
					break
				}
			}
			pid++
			continue
		} else {
			if pid+1 < len(p) {
				if p[pid+1] == '*' {
					if pid+2 < len(p) {
						pid += 2
						continue
					} else {
						return false
					}
				}
			} else {
				return false
			}
		}
	}
	if pid != len(p) {
		return false
	}
	if i < pid {
		return false
	}
	return true
}
