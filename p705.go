package projecteuler

// InversionCountSum calculates the sum of the inversion count for all possible
// divided sequences from the master sequence PrimeConcat(N), and returns the
// result module 10e9+7 Example: F(20) = 3312 and F(50) = 338079744. (where F
// is InversionCountSum)
// https://projecteuler.net/problem=705
func InversionCountSum(n int) int {
	primes := PrimeConcat(n)
	sequences := DividedSequence(primes)
	s := 0
	for _, seq := range sequences {
		s += InversionCount(seq)
	}
	return s%10e9 + 7
}

// InversionCount calculates the inversion count for a sequence.  The inversion
// count of a sequence of digits is the smallest number of adjacent pairs that
// must be swapped to sort the sequence.  For example, 34214 has inversion
// count of 5: 34214 -> 32414 -> 23414 -> 23144 -> 21344 -> 12344.
func InversionCount(s string) int {
	return 0
}

// DividedSequence generates a divided sequence.
// If each digit of a sequence is replaced by one of its divisors a divided
// sequence is obtained.  For example, the sequence 332 has 8 divided
// sequences: {332, 331, 312, 311, 132, 131, 112, 111}.
func DividedSequence(s string) []string {
	return nil
}

// PrimeConcat calculates to be the concatenation of all primes less than n,
// ignoring any zero digit.  For example, PrimeConcat(20) = 235711131719.
func PrimeConcat(n int) string {
	return ""
}
