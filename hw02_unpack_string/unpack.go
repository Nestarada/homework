package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"

	"golang.org/x/example/stringutil"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inp string) (res string, err error) {
	if len(inp) == 0 {
		return res, err
	} // input is empty
	if unicode.IsDigit(rune(inp[0])) {
		return "", ErrInvalidString
	} // input has first digit symbol

	repeats := 1
	for i := len(inp) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(inp[i])) && unicode.IsDigit(rune(inp[i+1])) {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(rune(inp[i])) {
			repeats, err = strconv.Atoi(string(inp[i]))
			if err != nil {
				return "", err
			}
			continue
		}

		for j := 0; j < repeats; j++ {
			res += string(inp[i])
		}
		repeats = 1
	}
	return stringutil.Reverse(res), nil
}
