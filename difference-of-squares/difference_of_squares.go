package diffsquares

func internal(n int) (x, y int) {
	for i := 1; i <= n; i++ {
		x += i
		y += i * i
	}
	return x * x, y
}

//SquareOfSum calculates the square of the first n natural numbers
func SquareOfSum(n int) (sum int) {
	result, _ := internal(n)
	return result
}

//SumOfSquares calculates sum of the squares of the first n natural number
func SumOfSquares(n int) (sum int) {
	_, result := internal(n)
	return result
}

//Difference calculates the difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	x, y := internal(n)
	return x - y
}
