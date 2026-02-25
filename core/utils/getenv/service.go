package getenv

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	logs "github.com/sirupsen/logrus"
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
		logs.Fatalf("[!] unsupported type: %T", any(*new(T)))
	}

	return result.(T)
}

func LoadEnv(filepath string) error {
	if err := godotenv.Load(".env"); err != nil {
		logs.Fatalf("[!] Error Load Env: %+v", err)
		return err
	}

	return nil
}
