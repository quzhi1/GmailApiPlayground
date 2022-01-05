package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient() *http.Client {
	b, err := os.ReadFile("accessToken")
	if err != nil {
		panic(err)
	}
	accessToken := string(b)
	fmt.Printf("Access token: %s", accessToken)
	fmt.Println()
	config := oauth2.Config{
		Endpoint: google.Endpoint,
	}
	token := oauth2.Token{
		AccessToken: accessToken,
		TokenType:   "Bearer",
	}
	return config.Client(context.Background(), &token)
}

func main() {
	ctx := context.Background()
	client := getClient()

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"
	r, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}
	if len(r.Labels) == 0 {
		fmt.Println("No labels found.")
		return
	}
	fmt.Println("Labels:")
	for _, l := range r.Labels {
		fmt.Printf("- %s\n", l.Name)
	}

	listRes, err := srv.Users.Messages.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve messages: %v", err)
	}
	if len(listRes.Messages) == 0 {
		fmt.Println("No messages found.")
		return
	}
	fmt.Println("Messages:")
	for _, l := range listRes.Messages {
		fmt.Printf("- %s\n", l.Id)
	}
}
