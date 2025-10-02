package phonenumber

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	var sb strings.Builder
	err := errors.New("")
	for _, c := range phoneNumber {
		if unicode.IsLetter(c) || strings.ContainsRune("@!:", c) {
			return "", err
		}
		if unicode.IsDigit(c) {
			sb.WriteRune(c)
		}
	}

	number := sb.String()
	l := len(number)

	if l < 10 || l > 11 {
		return "", err
	}
	if l == 11 {
		if !strings.HasPrefix(number, "1") {
			return "", err
		}
		number = number[1:]
	}
	afd := string(number[0])
	efd := string(number[3])
	if afd == "0" || afd == "1" {
		return "", err
	}
	if efd == "0" || efd == "1" {
		return "", err
	}

	return number, nil

}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	n, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", n[0:3], n[3:6], n[6:10]), nil
}
