package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
)

func main(){
	fmt.Println("Lets begin")
	godotenv.Load()
	portString := os.Getenv("PORT")

	if portString == ""{
		log.Fatal("PORT is not found in the env") 
	}
	router:= chi.NewRouter()
	fmt.Println("Port: ",portString)

}