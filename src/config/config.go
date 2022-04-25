package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConnectDB = ""
	Port            = 0
	SecretKey       []byte
)

func Load() {
	var error error
	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}
	Port, error = strconv.Atoi(os.Getenv("API_PORT"))
	if error != nil {
		Port = 9000
	}

	StringConnectDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
