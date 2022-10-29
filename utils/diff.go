package utils

import (
	"math"

	"github.com/shopspring/decimal"
)

func SwitchDiff(bigNumber string) float64 {
	bigDiff, _ := decimal.NewFromString(bigNumber)
	f, _ := decimal.NewFromInt(2).Pow(decimal.NewFromInt(256)).Sub(bigDiff).Float64()
	return 256 - math.Log2(f)
}
