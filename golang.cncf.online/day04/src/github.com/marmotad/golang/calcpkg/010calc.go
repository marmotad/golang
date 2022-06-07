package calcpkg

func Add(a, b int) int {
	return a + b
}

func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n*n - 1
}
