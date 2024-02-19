package main

import (
	"net/http"
	"os"

	"github.com/danielelegbe/go-chat/database"
	"github.com/danielelegbe/go-chat/handlers"
	"github.com/danielelegbe/go-chat/models"
	"github.com/danielelegbe/go-chat/ws"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	database.ConnectToDB()
	database.DB.AutoMigrate(&models.Message{})
}

func main() {
	mux := http.NewServeMux()

	room := ws.NewRoom()

	mux.HandleFunc("POST /messages", handlers.CreateMessage)
	mux.HandleFunc("GET /messages", handlers.GetMessages)
	mux.Handle("/subscribe", room)

	go room.Run()

	http.ListenAndServe(":8080", corsMiddleware(mux))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("HOST"))
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Check if the request is for the OPTIONS method ("preflight" CORS request)
		if r.Method == "OPTIONS" {
			// Respond with 200 OK without further processing
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler in line
		next.ServeHTTP(w, r)
	})
}
