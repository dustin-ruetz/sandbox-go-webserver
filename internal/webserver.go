package internal

import (

	// Excerpts from `log` package description:
	// > That logger writes to standard error and prints the date and time of each logged message.
	// > Every log message is output on a separate line: if the message being printed does not end in a newline, the logger will add one.
	"log"
	"net/http"

	"github.com/dustin-ruetz/sandbox-go-webserver/internal/handlers"
)

func StartWebserver() {
	http.HandleFunc("/api/", handlers.API)

	const port = ":4444"
	const url = "http://localhost" + port

	log.Println("Starting webserver...")
	log.Println("Routes:")
	log.Printf("  %s/api/", url)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("http.ListenAndServe: Failed to start webserver.", "\n err: ", err)
	}
}
