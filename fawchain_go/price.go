package main

import (
	"crypto/sha256"
	"encoding/binary"
	"time"
)

// determinePrice calculates the price based on the given seed and base price.
func determinePrice(seed string, basePrice float64) float64 {
	seedHash := sha256.Sum256([]byte(seed))
	currentTime := time.Now().Format("2006-01-02 15:04")
	timeHash := sha256.Sum256([]byte(currentTime))
	finalHashInput := append(seedHash[:], timeHash[:]...)
	finalHash := sha256.Sum256(finalHashInput)
	multiplier := float64(binary.BigEndian.Uint64(finalHash[0:8])) / float64(1<<64)
	multiplier = 0.001 + (10.999 * multiplier)
	return basePrice * multiplier
}
