package helpers

import (
	"math"
	"math/big"
)

func TruncateFloat64(f float64, unit float64) float64 {
	bf := big.NewFloat(0).SetPrec(1000).SetFloat64(f)
	bu := big.NewFloat(0).SetPrec(1000).SetFloat64(unit)

	bf.Quo(bf, bu)

	// Truncate:
	i := big.NewInt(0)
	bf.Int(i)
	bf.SetInt(i)

	f, _ = bf.Mul(bf, bu).Float64()

	return f
}

func TruncateFloat64Reward(n float64) float64 {
	decimals := 8

	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}
