package utils

import "strconv"

// Will return i on error
func AtoI(s string, errI int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return errI
	}

	return i
}
