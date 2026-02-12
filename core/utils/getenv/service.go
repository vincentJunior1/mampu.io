package getenv

import (
	"os"
	"strconv"
)

func Getenv[T any](name string) T {
	value := os.Getenv(name)

	var result any
	switch any(*new(T)).(type) {
	case string:
		result = value
	case int:
		v, _ := strconv.Atoi(value)
		result = v
	case bool:
		v, _ := strconv.ParseBool(value)
		result = v
	case float64:
		v, _ := strconv.ParseFloat(value, 64)
		result = v
	default:
		panic("unsupported type")
	}

	return result.(T)
}
