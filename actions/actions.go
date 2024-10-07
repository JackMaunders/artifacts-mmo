package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"jackmaunders/artifacts-mmo/request"
)

type Character struct {
	Name string `json:"name"`
}

type MoveCorrdinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c Character) GetCurrentLocation() MoveCorrdinates {
	fmt.Println("Getting " + c.Name + "'s current location")

	url := fmt.Sprintf("https://api.artifactsmmo.com/characters/%s", c.Name)

	responseBody := request.SendRequest(url, "GET", bytes.NewReader([]byte("")))
	response := request.HandleCharacterResponse(responseBody)

	currentLocation := MoveCorrdinates{X: response.Data.X, Y: response.Data.Y}

	fmt.Printf("%s's current location is %b %b", c.Name, currentLocation.X, currentLocation.Y)

	return currentLocation
}

func (c Character) Move(coordinates MoveCorrdinates) {
	fmt.Println(c.Name + " is moving to " + fmt.Sprint(coordinates))

	requestBody, _ := json.MarshalIndent(coordinates, "", "  ")

	url := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/move", c.Name)

	responseBody := request.SendRequest(url, "POST", bytes.NewReader(requestBody))
	response := request.HandleActionResponse(responseBody)

	fmt.Println(response.Data.Character.Name + " has moved to " + fmt.Sprint(response.Data.Destination.X) + " " + fmt.Sprint(response.Data.Destination.Y))
}
