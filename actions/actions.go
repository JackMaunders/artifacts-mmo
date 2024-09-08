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

func (c Character) Move(coordinates MoveCorrdinates) {
	fmt.Println(c.Name + " is moving to " + fmt.Sprint(coordinates))
	out, _ := json.MarshalIndent(coordinates, "", "  ")
	jsonBody := bytes.NewReader(out)

	url := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/move", c.Name)

	request.SendRequest(url, jsonBody)
}
