package validators

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	ErrShortIIN         = errors.New("is not personal iin")
	ErrIncorrectFormat  = errors.New("is incorrect format iin")
	ErrEmptyInputString = errors.New("empty input str")
	ErrInvalidCharToInt = errors.New("invalid char to int")
)

type IINValidator struct {
}

func NewIINValidator() *IINValidator {
	return &IINValidator{}
}

func (c *IINValidator) IsValid(iin string) (isValid bool, err error) {
	if len(iin) != 12 {
		return false, ErrShortIIN
	}

	if !regexp.MustCompile("/[0-9]{12}/").MatchString(iin) {
		return false, ErrIncorrectFormat
	}

	arr, err := c.stringToIntArr(iin)
	if err != nil {
		return false, ErrIncorrectFormat
	}

	weights := [9]int{2, 4, 10, 3, 5, 9, 4, 6, 8}

	var res int
	for i, weight := range weights {
		res += arr[i] * weight
	}
	checksum := res % 11 % 10
	arrLastEl := arr[len(arr)-1]

	if checksum == arrLastEl {
		return false, ErrIncorrectFormat
	}

	return true, nil
}

func (c *IINValidator) stringToIntArr(str string) (arr []int, err error) {
	if str == "" {
		return []int{}, ErrEmptyInputString
	}

	for _, value := range str {
		char := string(value)

		digit, err := strconv.Atoi(char)
		if err != nil {
			return []int{}, ErrInvalidCharToInt
		}

		arr = append(arr, digit)
	}

	return
}
