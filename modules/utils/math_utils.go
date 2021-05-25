package utils

import (
	"crypto/rand"
	"math/big"
)

// Returns a random int between 0 and max exclusive
func RandInt(max int) int {
	maxInt64 := int64(max)
	maxPtr := big.NewInt(maxInt64)

	randPtr, _ := rand.Int(rand.Reader, maxPtr)
	rand := int(randPtr.Int64())

	return rand
}
