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
	iin string
}

func NewIINValidator(iin string) *IINValidator {
	return &IINValidator{
		iin: iin,
	}
}

func (c *IINValidator) IsValid() (isValid bool, err error) {
	if len(c.iin) != 12 {
		return false, ErrShortIIN
	}

	if !regexp.MustCompile("/[0-9]{12}/").MatchString(c.iin) {
		return false, ErrIncorrectFormat
	}

	arr, err := c.stringToIntArr(c.iin)
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
		return true, ErrIncorrectFormat
	}

	return false, ErrIncorrectFormat
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
