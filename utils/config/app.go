package config

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

var AppEnv EnvSetup

type EnvSetup struct {
	UrlSvcTransaction string
	UrlSvcShipment    string
	JwtMayang         string
	UrlWeborder       string
}

func (e *EnvSetup) setEnvSetup() {
	e.UrlWeborder = os.Getenv("WEB_ORDER_URL")
}

func SetConfig() {
	LoadEnvVars()

	redis := new(RedisConfig)
	redis.SetConfigRedis().ConnectRedis()
	InitDb()
}

func LoadEnvVars() {
	cwd, _ := os.Getwd()
	dirString := strings.Split(cwd, "go-hexagonal")
	dir := strings.Join([]string{dirString[0], "go-hexagonal"}, "")
	AppPath := dir

	godotenv.Load(filepath.Join(AppPath, "/.env"))
	AppEnv.setEnvSetup()
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
