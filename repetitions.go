package toss

func repetitions(p int, n int) int {
	min, max := p, p

	// 最小
	if min > n {
		min = n
	}

	// 最大
	if max < n {
		max = n
	}

	// 期望最少次数
	switch {
	case min > 0 && max >= 5:
		return 3
	case min > 0 && max == 4:
		switch min {
		case 4:
			return 4
		case 3:
			return 4
		case 2:
			return 3
		case 1:
			return 3
		}
	case min > 0 && max == 3:
		switch min {
		case 3:
			return 3
		case 2:
			return 3
		case 1:
			return 2
		}
	case min > 0 && max == 2:
		switch min {
		case 2:
			return 2
		case 1:
			return 3
		}
	case min > 0 && max == 1:
		return 3
	default:
		return 6
	}

	return 99
}
