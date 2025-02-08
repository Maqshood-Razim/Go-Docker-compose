package main

import (
	db "docker-compose-go/config"
	"fmt"
	"log"
	"net/http"
)

func main() {

	database, err := db.ConnectDatabase()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	fmt.Println("Connected to Database:", database)

	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("worked"))
	})

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
