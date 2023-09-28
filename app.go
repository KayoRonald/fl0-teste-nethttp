package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Ok server",
		"status":  "ok",
	})
}

func main() {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		switch r.Method {
		case "GET":
			handleGetRequest(w, r)
		default:
			// Handle other HTTP methods
			fmt.Println("Handling other request")
		}
	})
	fmt.Printf("Rodando na porta %v\n", port)
	http.ListenAndServe(":"+port, nil)
}
