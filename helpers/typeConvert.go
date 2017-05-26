package helpers

import "strconv"
import "math"

// IsFloat is a helepr for testing if string is valid float64
func IsFloat(v string) bool {
	_, err := strconv.ParseFloat(v, 64)
	return err == nil
}

// IsInt is a helepr for testing if string is valid int32
func IsInt(v string) bool {
	_, err := strconv.ParseInt(v, 10, 32)
	return err == nil
}

// AsFloat assuming value is valid flaot64 returns corresponding float64 value of a string
func AsFloat(v string) float64 {
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

// AsInt assuming value is valid int32 returns corresponding int32 value of a string
func AsInt(v string) int {
	i, _ := strconv.ParseInt(v, 10, 32)
	return int(i)
}

// RadToDegrees returns degrees by redians
func RadToDegrees(rad float64) float64 {
	return rad * 180.0 / math.Pi
}
