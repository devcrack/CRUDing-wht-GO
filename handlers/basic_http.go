package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Structure for rpresent the JSON response

type Message struct {
	Content string    `json:"message"`
	Time    time.Time `json:"time"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	/*
		- w http.ResponseWriter: Un objeto donde escribir la respuesta HTTP
		- r *http.Request: Un puntero a un objeto que contiene la información de la solicitud HTTP
	*/
	// Comprueba si el método de la solicitud no es GET
	if r.Method != http.MethodGet {
		// envía un error HTTP 405 (Method Not Allowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Escribe el texto "Hello, World!" en la respuesta
	log.Println("Hola mundo desde simple get")
	fmt.Fprintf(w, "Hello Mundo")
}

// HelloJSONHandler maneja las peticiones al endpoint /hello/json
func HelloJSONHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Declaracion de la respuesta
	answer := Message{
		Content: "Hello Mundo",
		Time:    time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answer)
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	message := r.URL.Query().Get("message")
	if message == "" {
		message = "No message was provided"
	}
	fmt.Fprintf(w, "Echo: %s", message)
}
