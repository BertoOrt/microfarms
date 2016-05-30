package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

// Provider is the implementation of `goth.Provider` for accessing Google+.
type Provider struct {
	ClientKey   string
	Secret      string
	CallbackURL string
	config      *oauth2.Config
	prompt      oauth2.AuthCodeOption
}

// Gplus oauth json
type Gplus struct {
	ClientID                string   `json:"client_id"`
	ProjectID               string   `json:"project_id"`
	AuthURI                 string   `json:"auth_uri"`
	TokenURI                string   `json:"token_uri"`
	ClientSecret            string   `json:"client_secret"`
	AuthProviderX509CertURL string   `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string   `json:"clien_x509_cert_url"`
	RedirectURIS            []string `json:"redirect_uris"`
}

// Token from oauth json
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	IDToken     string `json:"id_token"`
}

// User information from oauth
type User struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Link    string `json:"link"`
	Picture string `json:"picture"`
}

var provider *Provider

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	file, err := ioutil.ReadFile("./gplus.json")
	if err != nil {
		log.Fatal(err)
	}
	var gplus Gplus
	json.Unmarshal(file, &gplus)
	// os.Getenv("TOKEN_SECRET") I still don't know when to jwt
	provider = newProvider(gplus, "http://localhost:8080/auth/callback")
}

func getAuthURL() string {
	var opts []oauth2.AuthCodeOption
	if provider.prompt != nil {
		opts = append(opts, provider.prompt)
	}
	url := provider.config.AuthCodeURL("state", opts...)
	return url
}

func newProvider(g Gplus, callbackURL string, scopes ...string) *Provider {
	p := &Provider{
		ClientKey:   g.ClientID,
		Secret:      g.ClientSecret,
		CallbackURL: callbackURL,
	}
	p.config = newConfig(p, g, scopes)
	return p
}

func newConfig(provider *Provider, g Gplus, scopes []string) *oauth2.Config {
	c := &oauth2.Config{
		ClientID:     provider.ClientKey,
		ClientSecret: provider.Secret,
		RedirectURL:  provider.CallbackURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  g.AuthURI,
			TokenURL: g.TokenURI,
		},
		Scopes: []string{},
	}
	c.Scopes = []string{"profile", "email", "openid"}
	return c
}

func fetchUser(code string) User {
	token := fetchToken(code)
	endpointProfile := "https://www.googleapis.com/oauth2/v2/userinfo"
	googleURL := endpointProfile + "?access_token=" + url.QueryEscape(token.AccessToken)
	fmt.Printf("url: %v", googleURL)
	response, err := http.Get(googleURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	bits, _ := ioutil.ReadAll(response.Body)
	var user User
	json.NewDecoder(bytes.NewReader(bits)).Decode(&user)
	return user
}

func fetchToken(code string) Token {
	params := "code=" + code + "&client_id=" + provider.ClientKey + "&client_secret=" + provider.config.ClientSecret + "&redirect_uri=" + provider.CallbackURL + "&grant_type=authorization_code"
	client := &http.Client{}
	req, _ := http.NewRequest("POST", provider.config.Endpoint.TokenURL, bytes.NewBufferString(params))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var token Token
	json.Unmarshal(body, &token)
	return token
}
