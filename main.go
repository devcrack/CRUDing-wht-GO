package main

import (
	"log"
	"net/http"
	"payroll_calculator/handlers"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	// Configuring routes
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/hello/json", handlers.HelloJSONHandler)
	http.HandleFunc("/echo", handlers.EchoHandler)

	port := ":8080"
	server := &http.Server{
		Addr:         port,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	log.Println("Starting server on port " + port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
