package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", server.CreateUser).Methods("POST")
	router.HandleFunc("/usuarios", server.SearchUsers).Methods("GET")
	router.HandleFunc("/usuarios/{id}", server.SearchUser).Methods("GET")
	router.HandleFunc("/usuarios/{id}", server.UpdateUser).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", server.DeleteUser).Methods("DELETE")

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

	// defer db.Close()
}
