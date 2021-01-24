// Package p705 addresses problem 705.
// https://projecteuler.net/problem=705
package p705

import (
	"math"
)

// InversionCountSum calculates the sum of the inversion count for all possible
// divided sequences from the master sequence PrimeConcat(N), and returns the
// result module 10e8+7 Example: F(20) = 3312 and F(50) = 338079744. (where F
// is InversionCountSum)
func InversionCountSum(n int) int {
	primes := PrimeConcat(n)
	s := DividedSequence(primes, InversionCount)
	return s % (1_000_000_007)
}

// ss is a central allocation for the InversionCount sequence.
// since InversionCount is not concurrent, there's no need to allocate
// a new slice for every call. It can just be reused.
var ss = []byte{}

// InversionCount calculates the inversion count for a sequence.  The inversion
// count of a sequence of digits is the smallest number of adjacent pairs that
// must be swapped to sort the sequence.  For example, 34214 has inversion
// count of 5: 34214 -> 32414 -> 23414 -> 23144 -> 21344 -> 12344.
func InversionCount(sequence []byte) int {
	if len(ss) != len(sequence) {
		ss = make([]byte, len(sequence))
	}
	copy(ss, sequence)

	n := len(ss)
	swaps := 0
	for {
		swapped := false
		for i := 1; i < n; i++ {
			// we can get away with comparing the bytes since we're only
			// comparing chars [0-9]
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

var divisors = map[byte][]byte{
	'9': {'1', '3', '9'},
	'8': {'1', '2', '4', '8'},
	'7': {'1', '7'},
	'6': {'1', '2', '3', '6'},
	'5': {'1', '5'},
	'4': {'1', '2', '4'},
	'3': {'1', '3'},
	'2': {'1', '2'},
	'1': {'1'},
}

// DividedSequence generates a divided sequence.
// If each digit of a sequence is replaced by one of its divisors a divided
// sequence is obtained.  For example, the sequence 332 has 8 divided
// sequences: {332, 331, 312, 311, 132, 131, 112, 111}.
//
// Note that the example (332...) above is kindof lacking. A better example is
// 432, which has 12 divided sequences (results shown in Test_DividedSequence).
func DividedSequence(s string, inversionCounter func([]byte) int) int {
	sequence := make([]byte, len(s))
	for i := 0; i < len(sequence); i++ {
		sequence[i] = '1'
	}

	// note, we're skipping a call to inversionCounter for the base value
	// of '111...111' since we know there are no inversions for that position.
	inversionCountTotal := 0
	for i := len(sequence) - 1; i >= 0; {
		incValue, carry := increment(sequence[i], s[i])
		sequence[i] = incValue
		if carry {
			i--
		} else {
			inversionCountTotal += inversionCounter(sequence)
			i = len(sequence) - 1
		}
	}
	return inversionCountTotal
}

// increment returns the next value and a bool representing whether or not a
// carry occured
func increment(digit, base byte) (byte, bool) {
	divs := divisors[base]
	for i := 0; i < len(divs); i++ {
		if divs[i] == digit {
			if i == len(divs)-1 {
				return divs[0], true
			}
			return divs[i+1], false
		}
	}
	panic("digit not found for base")
}

// PrimeConcat calculates the concatenation of all primes less than n, ignoring
// any zero digit.  For example, PrimeConcat(20) = 235711131719.
func PrimeConcat(m int) string {
	s := ""
	for n := 1; n < m; n++ {
		if isPrime(n) {
			_ = n
			// main questions:
			//  how many primes are there below 10e8?
			//		~5761455
			//  how many bytes to store those primes as int32s?
			//		~23MB
			//	how many bytes to store those as a byte (char) array?
			//		<64MB
			// how long to calculate them?
			//		~15 minutes?
			// _ = strconv.Itoa(n)
			// s = s + strconv.Itoa(n)
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
