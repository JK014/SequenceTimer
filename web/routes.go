package web

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// RegisterRoutes defines all application routes
func RegisterRoutes(r *mux.Router) {
	// Stop handler
	r.HandleFunc("/stop", StopHandler).Methods("POST")

}

func StopHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Shutting down the server..."))
	if err != nil {
		log.Printf("Error writing response: %v\n", err)
	}
	close(stopChannel) // Notify the server to shut down
}
