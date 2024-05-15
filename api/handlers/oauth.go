package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type OAuthHandler struct{}

func NewOAuthHandler() *OAuthHandler {
	return &OAuthHandler{}
}

func generateState() (string, error) {
	// Generate a random 8-byte state
	state := make([]byte, 8)
	if _, err := rand.Read(state); err != nil {
		return "", err
	}
	return hex.EncodeToString(state), nil
}

func (h *OAuthHandler) AuthWhoop(c *gin.Context) {
	clientID := os.Getenv("WHOOP_CLIENT_ID")
	redirectURI := os.Getenv("REDIRECT_URI")
	scope := "read:recovery read:cycles read:workout read:sleep read:profile read:body_measurement"

	state, err := generateState()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate state"})
		return
	}

	authURL := fmt.Sprintf("https://api.prod.whoop.com/oauth/oauth2/auth?client_id=%s&response_type=code&redirect_uri=%s&scope=%s&state=%s", clientID, redirectURI, scope, state)
	c.Redirect(http.StatusFound, authURL)
}

func (h *OAuthHandler) Callback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	clientID := os.Getenv("WHOOP_CLIENT_ID")
	clientSecret := os.Getenv("WHOOP_CLIENT_SECRET")
	redirectURI := os.Getenv("REDIRECT_URI")

	// Todo: validate the state parameter to prevent CSRF attacks
	if state == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "State parameter is missing"})
		return
	}

	tokenURL := "https://api.prod.whoop.com/oauth/oauth2/token"
	reqBody := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=authorization_code&code=%s&redirect_uri=%s", clientID, clientSecret, code, redirectURI)
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(reqBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to request token"})
		return
	}
	defer resp.Body.Close()

	var tokenResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode token response"})
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}
