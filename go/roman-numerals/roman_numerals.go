package romannumerals

import (
	"errors"
	"strings"
)

func DigitToRomanNumeral(n int, smallDigit int) string {
	var small, aux, large string
	switch smallDigit {
	case 1000:
		small = "M"
	case 100:
		small = "C"
		aux = "D"
		large = "M"
	case 10:
		small = "X"
		aux = "L"
		large = "C"
	case 1:
		small = "I"
		aux = "V"
		large = "X"
	}

	if n <= 3 {
		return strings.Repeat(small, n)
	}
	if n == 4 {
		return small + aux
	}
	if n == 5 {
		return aux
	}
	if n > 5 && n < 9 {
		return aux + strings.Repeat(small, n-5)
	}
	if n == 9 {
		return small + large
	}
	return ""
}

func ToRomanNumeral(n int) (string, error) {
	if n > 3000 {
		return "", errors.New("Input number must not be greater than 3000")
	}
	if n < 1 {
		return "", errors.New("Input must be greater than zero.")
	}
	var output string
	for p := 1000; p >= 1; p = p / 10 {
		digit := n / p
		n = n % p
		if digit != 0 {
			output += DigitToRomanNumeral(digit, p)
		}
	}

	return output, nil
}
