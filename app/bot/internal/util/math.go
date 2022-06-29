package util

func Abs(n int64) int64 {
	if n >= 0 {
		return n
	}
	return -n
}
