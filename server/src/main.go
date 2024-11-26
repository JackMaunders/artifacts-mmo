package main

import (
	"fmt"
	"jackmaunders/artifacts-mmo/src/actions"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file, using CLI environment variables\n")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Panic("Port not configured, exiting\n")
	}

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/{character}/move/{direction}", handleMove).Methods(http.MethodPost, http.MethodOptions)

	fmt.Printf("Server is starting on port %s, head to https://artifactsmmo.com/client to see the game in action\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

const (
	Up    = "up"
	Down  = "down"
	Left  = "left"
	Right = "right"
)

func handleMove(w http.ResponseWriter, r *http.Request) {
	// @TODO: Add CORS middleware & handle set origins
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	characterName := vars["character"]
	direction := vars["direction"]

	// Move to regex validation in path handler
	if !isValidDirection(direction) {
		http.Error(w, "Invalid direction provided", http.StatusBadRequest)
		return
	}

	character := actions.Character{Name: characterName}
	currentCoordinates := character.GetCurrentLocation()
	character.Move(getNewCoordinates(direction, currentCoordinates))

	w.WriteHeader(http.StatusOK)
}

func isValidDirection(direction string) bool {
	switch direction {
	case Up, Down, Left, Right:
		return true
	default:
		return false
	}
}

func getNewCoordinates(direction string, currentCoordinates actions.MoveCorrdinates) actions.MoveCorrdinates {
	switch direction {
	case Up:
		return actions.MoveCorrdinates{X: currentCoordinates.X, Y: currentCoordinates.Y - 1}
	case Down:
		return actions.MoveCorrdinates{X: currentCoordinates.X, Y: currentCoordinates.Y + 1}
	case Left:
		return actions.MoveCorrdinates{X: currentCoordinates.X - 1, Y: currentCoordinates.Y}
	case Right:
		return actions.MoveCorrdinates{X: currentCoordinates.X + 1, Y: currentCoordinates.Y}
	}

	return currentCoordinates
}
