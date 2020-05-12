package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("env/local.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routes := GetRoutes()

	fmt.Println(os.Getenv("PORT"))
	log.Fatal(routes.Run())

}
