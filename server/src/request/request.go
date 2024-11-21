package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"jackmaunders/artifacts-mmo/src/schema"
	"net/http"
	"os"
)

type ResponseData struct {
	Character   schema.Character   `json:"character"`
	Cooldown    schema.Cooldown    `json:"cooldown"`
	Destination schema.Destination `json:"destination"`
}

type ActionResponse struct {
	Data ResponseData `json:"data"`
}

type CharacterResponse struct {
	Data schema.Character `json:"data"`
}

type ResponseErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseError struct {
	Error ResponseErrorData `json:"error"`
}

func SendRequest(url string, method string, payload *bytes.Reader) []byte {

	fmt.Println("Sending request to:", url)

	req, _ := http.NewRequest(method, url, payload)

	req = addHeaders(req)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Error making request: %v", err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != 200 {
		var errorBody ResponseError
		json.Unmarshal(body, &errorBody)
		prettyError, _ := json.MarshalIndent(errorBody, "", "\t")
		fmt.Println(string(prettyError))
	}

	return body
}

func HandleActionResponse(body []byte) ActionResponse {
	var result ActionResponse
	json.Unmarshal(body, &result)
	return result
}

func HandleCharacterResponse(body []byte) CharacterResponse {
	var result CharacterResponse
	json.Unmarshal(body, &result)
	return result
}

func addHeaders(req *http.Request) *http.Request {
	artifactsMmoToken := getArtifactsMmoToken()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+artifactsMmoToken)
	return req
}

func getArtifactsMmoToken() string {
	return os.Getenv("ARTIFACTS_MMO_TOKEN")
}
