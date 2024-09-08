package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"jackmaunders/artifacts-mmo/env"
	"jackmaunders/artifacts-mmo/schema"
	"net/http"
)

type ResponseData struct {
	Character   schema.Character   `json:"character"`
	Cooldown    schema.Cooldown    `json:"cooldown"`
	Destination schema.Destination `json:"destination"`
}

type Response struct {
	Data ResponseData `json:"data"`
}

type ResponseErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseError struct {
	Error ResponseErrorData `json:"error"`
}

func SendRequest(url string, payload *bytes.Reader) error {
	artifactsMmoToken := getArtifactsMmoToken()

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+artifactsMmoToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != 200 {
		var errorBody ResponseError
		json.Unmarshal(body, &errorBody)
		prettyError, _ := json.MarshalIndent(errorBody, "", "\t")
		fmt.Printf(string(prettyError))
		return nil
	}

	var result Response
	json.Unmarshal(body, &result)
	pretty, _ := json.MarshalIndent(result, "", "\t")
	fmt.Println(res)
	fmt.Println(string(pretty))
	return nil
}

func getArtifactsMmoToken() string {
	return env.GetEnvVar("ARTIFACTS_MMO_TOKEN")
}
