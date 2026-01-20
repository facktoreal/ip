package helpers

import (
	"fmt"
	"math"
)

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateHash(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+"%s", val, suffix)

}

// Hashes produces a human readable representation of an SI size.
func Hashes(s uint64) string {
	sizes := []string{"H/s", "kH/s", "MH/s", "GH/s", "TH/s", "PH/s", "EH/s"}

	return humanateHash(uint64(s), 1000, sizes)
}
