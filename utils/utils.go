package utils

import (
	"fmt"
	"log"
	"os"
)

func MustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal(fmt.Errorf("%s environment variable must be set", key))
	}

	return val
}
