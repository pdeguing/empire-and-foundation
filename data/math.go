package data

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func ClipInt64(i, min, max int64) int64 {
	if i <= min {
		return min
	}
	if i >= max {
		return max
	}
	return i
}
