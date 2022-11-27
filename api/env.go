package api

import (
	"os"
)

// Env

func Rp_addENV(key, value string) error {

	err := os.Setenv(key, value)
	return err

}

func Rp_removeENV(key string) error {

	err := os.Unsetenv(key)

	return err
}

func Rp_getENV(key string) string {

	val := os.Getenv(key)

	return val

}

func Rp_clearENV() {

	os.Clearenv()

}
