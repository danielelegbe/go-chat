package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/danielelegbe/go-chat/database"
	"github.com/danielelegbe/go-chat/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func CreateMessage(w http.ResponseWriter, r *http.Request) {

	var message models.Message

	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	result := database.DB.Create(&message)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		ID      int    `json:"id"`
		Content string `json:"content"`
	}

	var messages []Response

	result := database.DB.Model(&models.Message{}).Find(&messages)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)

}
