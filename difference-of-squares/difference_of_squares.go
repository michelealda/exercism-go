package diffsquares

//SquareOfSum calculates the square of the first n natural numbers
func SquareOfSum(n int) (sum int) {
	gaussSum := (n * (n + 1)) / 2
	return gaussSum * gaussSum
}

//SumOfSquares calculates sum of the squares of the first n natural number
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

//Difference calculates the difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
