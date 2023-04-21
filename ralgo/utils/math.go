package utils

import (
	"bytes"
	"strconv"
)

func PrimeFactors(n int) []int64 {
	var factors []int64

	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, int64(i))
			n /= i
		}
	}

	return factors
}

func EuclideanDivision(dividend, divisor int) (quotient, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor

	return quotient, remainder
}

func CoupleToInt64(chunk []byte, couple KeyCouple) int64 {
	chunk = bytes.Replace(chunk, []byte(couple.A), []byte("0"), -1)
	chunk = bytes.Replace(chunk, []byte(couple.B), []byte("1"), -1)

	s, err := strconv.ParseInt(string(chunk), 2, 64)
	CheckError(err)

	return s
}
