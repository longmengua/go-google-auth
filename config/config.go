package config

import (
	"fmt"
	"go-google-auth/util"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Init() {
	path, err := util.GetRootPath()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	p := strings.Replace(path, "/cmd", "", -1)
	p = fmt.Sprintf("%s/.env", p)
	err = godotenv.Load(p)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
