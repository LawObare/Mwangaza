package utils

import "regexp"

func IsValidPhone(phone string) bool {
	matched, _ := regexp.MatchString(`^\+?254\d{9}$`, phone)
	return matched
}

func IsValidFarmID(id int) bool {
	return id > 0
}