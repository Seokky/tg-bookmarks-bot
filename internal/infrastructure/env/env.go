// Package env manages access to environment variables
package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

// loadOnce performs "once" version of [godotenv.Load]
func loadOnce() {
	sync.OnceFunc(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("can't load .env: ", err)
		}
	})()
}

// Get performs access to environment variable with automatically load env variables once
func Get(key string) string {
	loadOnce()

	return os.Getenv(key)
}
