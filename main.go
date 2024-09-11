package main

import (
	"jackmaunders/artifacts-mmo/actions"
	"net/http"
)

func handleMove(w http.ResponseWriter, r *http.Request) {
	jack := actions.Character{Name: "Jack"}
	jack.Move(actions.MoveCorrdinates{X: 5, Y: 4})
}

func main() {
	http.HandleFunc("/move", handleMove)

	err := http.ListenAndServe(":3001", nil)
}
