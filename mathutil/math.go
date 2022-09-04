package mathutil

// MaxInt64 ...
func MaxInt64(is ...int64) int64 {
	max := is[0]
	for _, v := range is {
		if v > max {
			max = v
		}
	}
	return max
}

// MinInt64 ...
func MinInt64(is ...int64) int64 {
	min := is[0]
	for _, v := range is {
		if v < min {
			min = v
		}
	}
	return min
}
