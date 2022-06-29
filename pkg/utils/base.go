package utils

func IF(f bool, a, b interface{}) interface{} {
	if f {
		return a
	}
	return b
}
