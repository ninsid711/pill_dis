package envloader

import (
	"github.com/joho/godotenv"
	"log"
)

func Loadenv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file :( ")
	}
}
