package config

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

func SetConfig() {
	LoadEnvVars()

	redis := new(RedisConfig)
	redis.SetConfigRedis().ConnectRedis()
}

func LoadEnvVars() {
	dir, _ := os.Getwd()
	AppPath := dir

	godotenv.Load(filepath.Join(AppPath, "/.env"))
}

// LoadEnvVarsLocal is load .env.local is used to load env vars for local development. For testing purposes
func LoadEnvVarsLocal() {
	dir := RootDir()
	AppPath := dir

	godotenv.Load(filepath.Join(AppPath, "/.env.local"))
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	// delete int string return /utils to get the root dir
	rootDir := strings.TrimRight(filepath.Dir(d), "/utils")
	return rootDir
}
