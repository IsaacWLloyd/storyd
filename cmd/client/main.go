package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseURL = "http://localhost:8080"
)

func main() {
	// Test Create Lobby
	lobbyID := createLobby()

	// Test Get Lobby
	getLobby(lobbyID)
}

func createLobby() string {
	url := baseURL + "/lobbies"
	requestBody, err := json.Marshal(map[string]string{
		"id": "test-lobby",
	})
	if err != nil {
		log.Fatalf("Error creating request body: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error creating lobby: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Error parsing response JSON: %v", err)
	}

	fmt.Printf("Create Lobby Response: %+v\n", response)
	return response["lobbyId"].(string)
}

func getLobby(lobbyID string) {
	url := fmt.Sprintf("%s/lobbies/%s", baseURL, lobbyID)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error getting lobby: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var lobby map[string]interface{}
	err = json.Unmarshal(body, &lobby)
	if err != nil {
		log.Fatalf("Error parsing response JSON: %v", err)
	}

	fmt.Printf("Get Lobby Response: %+v\n", lobby)
}