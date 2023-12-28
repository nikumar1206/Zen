package utils

import (
	"strconv"
)

// check if string is an integer
func IsInt(s string) bool {

	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
