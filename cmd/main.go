package main

import (
	"fmt"
	"go-google-auth/config"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	config.Init()

	// Replace with your Google OAuth credentials
	clientID := config.Get("GOOGLE_CLIENT_ID")
	clientSecret := config.Get("GOOGLE_CLIENT_SECRET")
	redirectURL := config.Get("GOOGLE_CALLBACK_URL")

	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	http.HandleFunc("/auth/google", func(w http.ResponseWriter, r *http.Request) {
		url := config.AuthCodeURL("state")
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/auth/google/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		token, err := config.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Fatal("Failed to exchange token:", err)
		}

		// Use the token to access Google API
		// For example, you can make requests to the Google+ API using the token

		fmt.Fprintf(w, "Token: %s\n", token.AccessToken)
	})

	log.Println("server up on port:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
