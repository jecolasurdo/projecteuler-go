package projecteuler

import (
	"math"
	"strconv"
)

// InversionCountSum calculates the sum of the inversion count for all possible
// divided sequences from the master sequence PrimeConcat(N), and returns the
// result module 10e8+7 Example: F(20) = 3312 and F(50) = 338079744. (where F
// is InversionCountSum)
// https://projecteuler.net/problem=705
//
// Areas to improve performance:
//	- InversionCount uses can use a more optimized bubble sort implementation
//  - InversionCount for each sequence can be calculated in parallel
//  - InversionCount can be initialized as soon as a unique sequence is found.
//		This might not really gain very much though.
//	- PrimeConcat can be calculated once and cached on disk between runs.
//	- Can probably look for areas to reduce heap usage
//	- Once an inversionCount for a sequence is calculated, it's not necessary
//		to store that sequence in memory any longer. The only need to keep
//		a sequence around is to ensure that DividedSequence can identify duplicates
//		efficiently. This can probably be done more efficiently through some means
//		other than a basic set. Rather than storing the sequences in the set,
//		hashes of each sequence can be stored. Or compressed versions can be stored.
//		Or, potentially, there is another data structure that would be more
//		efficient for the operation in general.
func InversionCountSum(n int) int {
	primes := PrimeConcat(n)
	sequences := DividedSequence(primes)
	s := 0
	for seq := range sequences {
		s += InversionCount(seq)
	}
	return s % (1_000_000_007)
}

// InversionCount calculates the inversion count for a sequence.  The inversion
// count of a sequence of digits is the smallest number of adjacent pairs that
// must be swapped to sort the sequence.  For example, 34214 has inversion
// count of 5: 34214 -> 32414 -> 23414 -> 23144 -> 21344 -> 12344.
func InversionCount(ss []byte) int {
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
func DividedSequence(s string) chan []byte {
	sequenceSource := make(chan []byte, 10)
	go func() {
		defer close(sequenceSource)
		sequence := make([]byte, len(s))
		for i := 0; i < len(sequence); i++ {
			sequence[i] = '1'
		}
		initialSequence := make([]byte, len(sequence))
		copy(initialSequence, sequence)
		sequenceSource <- initialSequence

		for i := len(sequence) - 1; i >= 0; {
			newSequence := make([]byte, len(sequence))
			incValue, carry := increment(sequence[i], s[i])
			sequence[i] = incValue
			if carry {
				i--
			} else {
				copy(newSequence, sequence)
				sequenceSource <- newSequence
				i = len(sequence) - 1
			}
		}
	}()
	return sequenceSource
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
