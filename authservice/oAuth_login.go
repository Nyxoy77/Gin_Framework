package authservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOAuthConfig = &oauth2.Config{
	ClientID:     "625471794559-pqj6ka5n4a8si32nrs6s9urgi0t58m5r.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  "https://localhost:8081/user/oauth/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

var oauthStateString = "random_string"

func HandleGoogleLogin(c *gin.Context) {
	url := googleOAuthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}
