package main

import (
	"log"

	"github.com/YanyChoi/fiber-poc/routers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	app := routers.NewRouter()
	app.Run()
}