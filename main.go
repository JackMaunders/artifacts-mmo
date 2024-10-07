package main

import (
	"fmt"
	"jackmaunders/artifacts-mmo/actions"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{character}/move/{direction}", handleMove).Methods("POST")

	fmt.Println("Server is starting on port 3001")
	http.ListenAndServe(":3001", r)
}

func handleMove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	characterName := vars["character"]
	direction := vars["direction"]

	fmt.Println("Character: ", characterName)
	fmt.Println("Direction: ", direction)

	// Move to regex validation in path handler
	if direction != "north" && direction != "south" && direction != "east" && direction != "west" {
		http.Error(w, "Invalid direction provided", http.StatusBadRequest)
		return
	}

	character := actions.Character{Name: characterName}
	currentCoordinates := character.GetCurrentLocation()
	character.Move(getNewCoordinates(direction, currentCoordinates))

	w.WriteHeader(http.StatusOK)
}

func getNewCoordinates(direction string, currentCoordinates actions.MoveCorrdinates) actions.MoveCorrdinates {
	switch direction {
	case "north":
		return actions.MoveCorrdinates{X: currentCoordinates.X, Y: currentCoordinates.Y - 1}
	case "south":
		return actions.MoveCorrdinates{X: currentCoordinates.X, Y: currentCoordinates.Y + 1}
	case "east":
		return actions.MoveCorrdinates{X: currentCoordinates.X + 1, Y: currentCoordinates.Y}
	case "west":
		return actions.MoveCorrdinates{X: currentCoordinates.X - 1, Y: currentCoordinates.Y}
	}

	return currentCoordinates
}
