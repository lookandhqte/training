package hexlet

func titleString(title string) bool {
	var ans bool
	//65-90 заглавные
	// 32 = " "
	for i := 0; i < len(title); i++ {
		if 91 > title[i] && title[i] < 64 {
			for j := i; j < len(title); j++ {
				if title[j] == 32 {
					i = j + 1
					break
				}
			}
		} else {
			ans = false
			break
		}
	}
	return ans
}
