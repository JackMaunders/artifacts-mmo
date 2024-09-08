package main

import (
	"jackmaunders/artifacts-mmo/actions"
)

func main() {
	jack := actions.Character{Name: "Jack"}
	jack.Move(actions.MoveCorrdinates{X: 5, Y: 4})
}
