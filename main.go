package main

import (
	"os"
	"log"
	"net/http"
	"tracker/handlers"
)

func main() {
	
	if len(os.Args) > 1 {
		log.Println("The program epects only one argument: \n \n e.g go run main.go ")
		return
	  }
	// Map URLs to their respective page handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/relations", handlers.RelationHandler)
	http.HandleFunc("/locations", handlers.LocationHandler)
	http.HandleFunc("/dates/", handlers.DatesHandler)
	http.HandleFunc("/artistProfile", handlers.ArtistDetails)
	http.HandleFunc("/artist", handlers.ArtistHandler)

	// Serve static files (CSS, images)
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Start the server
	log.Println("Server running at http://localhost:8087")
	err := http.ListenAndServe(":8087", nil)
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
