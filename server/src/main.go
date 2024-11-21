package main

import (
	"fmt"
	"jackmaunders/artifacts-mmo/src/actions"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{character}/move/{direction}", handleMove).Methods("POST")

	fmt.Println("Server is starting on port 3001, head to https://artifactsmmo.com/client to see the game in action")
	http.ListenAndServe(":3001", r)
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
