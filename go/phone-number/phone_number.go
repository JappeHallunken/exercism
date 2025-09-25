package phonenumber

import (
	"fmt"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	var digits []rune
	for _, r := range phoneNumber {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		}
	}

	length := len(digits)
	if length < 10 || length > 11 {
		return "", fmt.Errorf("not a valid number: %s", phoneNumber)
	}

	if length == 11 {
		if digits[0] != '1' || digits[1] < '2' || digits[4] < '2' {
			return "", fmt.Errorf("not a valid number: %s", phoneNumber)
		}
		digits = digits[1:]
	} else { // length == 10
		if digits[0] < '2' || digits[3] < '2' {
			return "", fmt.Errorf("not a valid number: %s", phoneNumber)
		}
	}

	return string(digits), nil
}

func AreaCode(phoneNumber string) (string, error) {
	areaCode, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return areaCode[:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}
