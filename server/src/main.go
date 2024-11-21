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
	r.HandleFunc("/{character}/move/{direction}", handleMove).Methods("POST")

	fmt.Printf("Server is starting on port %s, head to https://artifactsmmo.com/client to see the game in action\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

const (
	North = "north"
	South = "south"
	East  = "east"
	West  = "west"
)

func handleMove(w http.ResponseWriter, r *http.Request) {
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
	case North, South, East, West:
		return true
	default:
		return false
	}
}

func getNewCoordinates(direction string, currentCoordinates actions.MoveCorrdinates) actions.MoveCorrdinates {
	switch direction {
	case North:
		return actions.MoveCorrdinates{X: currentCoordinates.X, Y: currentCoordinates.Y - 1}
	case South:
		return actions.MoveCorrdinates{X: currentCoordinates.X, Y: currentCoordinates.Y + 1}
	case East:
		return actions.MoveCorrdinates{X: currentCoordinates.X + 1, Y: currentCoordinates.Y}
	case West:
		return actions.MoveCorrdinates{X: currentCoordinates.X - 1, Y: currentCoordinates.Y}
	}

	return currentCoordinates
}
