package main

import (
	"log"
	"net/http"
	"os"

	// "diawan.club/web_in_go2/base"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	log.Fatal(
		"[main] start program",
	)

	log.Fatal(http.ListenAndServe(":"+port, Router(mux))) // Porting listening

}
