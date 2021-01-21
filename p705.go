package projecteuler

import (
	"math"
	"strconv"
	"strings"
)

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
	ss := strings.Split(s, "")
	n := len(ss)
	swaps := 0
	for {
		swapped := false
		for i := 1; i < n-1; i++ {
			if ss[i-1] > ss[i] {
				ss[i-1], ss[i] = ss[i], ss[i-1]
				swaps++
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
	return swaps
}

// DividedSequence generates a divided sequence.
// If each digit of a sequence is replaced by one of its divisors a divided
// sequence is obtained.  For example, the sequence 332 has 8 divided
// sequences: {332, 331, 312, 311, 132, 131, 112, 111}.
func DividedSequence(s string) []string {
	sequence := []string{s}
	for i := 0; i < len(s); i++ {
		newSequence := []string{}
		for _, item := range sequence {
			if i == len(s) {
				newSequence = append(newSequence, item[:i]+"1")
			} else {
				newSequence = append(newSequence, item[:i]+"1"+item[i+1:])
			}
		}
		sequence = append(sequence, newSequence...)
	}
	return sequence
}

// PrimeConcat calculates the concatenation of all primes less than n, ignoring
// any zero digit.  For example, PrimeConcat(20) = 235711131719.
func PrimeConcat(m int) string {
	s := ""
	for n := 1; n < m; n++ {
		if isPrime(n) {
			s = s + strconv.Itoa(n)
		}
	}
	return s
}

func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; math.Pow(float64(i), 2.) <= float64(n); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
