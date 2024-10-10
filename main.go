package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

// Load environment variables from .env file
func loadEnv() {
	err := godotenv.Load("keystoken.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// Creating an OAuth1 client using user access tokens for Twitter API
func getOAuth1func() *http.Client {
	config := oauth1.NewConfig(os.Getenv("X_API_KEY"), os.Getenv("X_API_SECRET_KEY"))
	token := oauth1.NewToken(os.Getenv("X_ACCESS_TOKEN"), os.Getenv("X_ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	return httpClient
}

// Function to post a tweet using Twitter API
func postTweet(client *http.Client, tweetText string) {
	url := "https://api.twitter.com/2/tweets"

	// Create a JSON request body
	requestBody, err := json.Marshal(map[string]string{
		"text": tweetText,
	})
	if err != nil {
		log.Fatal("Error creating JSON request body:", err)
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Perform the POST request
	response, err := client.Do(req)
	if err != nil {
		log.Fatal("Error posting tweet:", err)
	}
	defer response.Body.Close()

	// Read and output the response for debugging
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	// Check if the request was successful
	if response.StatusCode != http.StatusCreated {
		log.Fatalf("Failed to post tweet. Status Code: %d. Response: %s", response.StatusCode, body)
	}

	fmt.Println("Tweet posted successfully!")
	fmt.Println(string(body))
}

// Function to delete a tweet using Twitter API
func deleteTweet(client *http.Client, tweetID string) {
	url := fmt.Sprintf("https://api.twitter.com/2/tweets/%s", tweetID)

	// Creating a new DELETE request
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Perform the DELETE request
	response, err := client.Do(req)
	if err != nil {
		log.Fatal("Error deleting tweet:", err)
	}
	defer response.Body.Close()

	// Read and output the response for debugging
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	// Check if the request was successful
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Failed to delete tweet. Status Code: %d. Response: %s", response.StatusCode, body)
	}

	fmt.Println("Tweet deleted successfully!")
	fmt.Println(string(body))
}

func main() {
	// Load environment variables from the .env file
	loadEnv()

	// Create a new Twitter API client using OAuth 1.0a
	client := getOAuth1func()

	// User input for selecting the action (post or delete)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What would you like to do? (post/delete): ")
	action, _ := reader.ReadString('\n')
	action = strings.TrimSpace(action)

	if action == "post" {
		// Post a tweet
		fmt.Println("Enter your tweet: ")
		tweetText, _ := reader.ReadString('\n')
		tweetText = strings.TrimSpace(tweetText)
		postTweet(client, tweetText)
	} else if action == "delete" {
		// Delete a tweet
		fmt.Println("Enter the tweet ID to delete: ")
		tweetID, _ := reader.ReadString('\n')
		tweetID = strings.TrimSpace(tweetID)
		deleteTweet(client, tweetID)
	} else {
		fmt.Println("Invalid action. Please enter 'post' or 'delete'.")
	}
}
