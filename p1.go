package projecteuler

// FindSumOfMultiples solves the following problem: If we list all the natural
// numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum
// of these multiples is 23.  Find the sum of all the multiples of 3 or 5 below
// 1000.

// FindSumOfMultiplesBruteForce .
func FindSumOfMultiplesBruteForce(baseFactors []int, sumLimit int) int {
	s := 0
	for n := 1; n < sumLimit; n++ {
		for _, factor := range baseFactors {
			if n%factor == 0 {
				s += n
				break
			}
		}
	}

	return s
}

// FindSumOfMultiples .
func FindSumOfMultiples(baseFactors []int, sumLimit int) int {
	multiples := map[int]struct{}{}
	sum := 0
	for _, baseFactor := range baseFactors {
		for n := 1; n*baseFactor < sumLimit; n++ {
			multiple := n * baseFactor
			if _, found := multiples[multiple]; !found {
				sum += multiple
				multiples[multiple] = struct{}{}
			}
		}
	}
	return sum
}
