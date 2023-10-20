package phonenumber

import (
	"regexp"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	re := regexp.MustCompile("[0-9]").FindAllString(phoneNumber, -1)
	response := strings.Join(re, "")
	return response, nil
}
