package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Crear el enrutador
	router := mux.NewRouter()
	const port string = ":8000"

	//handlers o manejadores de rutas invocamos a la función HandleFunc
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		//Escribe en el ResponseWriter (Postman)
		fmt.Fprintln(resp, "Up and running")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}