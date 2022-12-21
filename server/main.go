package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Iniciou a request.")

	defer log.Println("Finalizou a request.")

	select {
	case <-time.After(time.Second * 5):
		fmt.Println(w, "Requisição processada com sucesso.")
		w.Write([]byte("Requisição processada com sucesso.\n"))
	case <-ctx.Done():
		log.Println("Request cancelada.")
		http.Error(w, "Request cancelada", http.StatusRequestTimeout)
	}
}
