package utils

type math struct{}

var Math = math{}

func (m math) MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
