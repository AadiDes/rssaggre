package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code> 499{ //400 range codes are client side, ignore
		log.Println("Responding with 5XX error: ", msg)
	}

	type errResponse struct{
		Error string `json:"error"`
	}
	respondWithJson(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}){
	dat, er := json.Marshal(payload) //json string returned as bytes
	if er !=nil{
		log.Printf("Failed to marshal json respons: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)

}
