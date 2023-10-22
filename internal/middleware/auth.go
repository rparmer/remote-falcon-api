package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rparmer/remote-falcon-api/internal/database"
	"github.com/rparmer/remote-falcon-api/internal/service/auth"
	"golang.org/x/crypto/bcrypt"
)

const userkey = "user"

var authService *auth.Service

func Configure(db *database.MongoDB) {
	authService = auth.New(db)
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

func PluginAuthRequired(c *gin.Context) {
	if checkBasicAuth(c) || checkTokenAuth(c) {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if username != "admin" || password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set(userkey, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func checkBasicAuth(c *gin.Context) bool {
	authHeader := c.GetHeader("Authorization")
	splitCreds := strings.Split(authHeader, "Basic ")
	if len(splitCreds) != 2 {
		return false
	}
	creds := splitCreds[1]
	rawCreds, _ := base64.StdEncoding.DecodeString(creds)
	splitRawCreds := strings.Split(string(rawCreds), ":")
	if len(splitRawCreds) != 2 {
		return false
	}
	username := splitRawCreds[0]
	password := splitRawCreds[1]
	auth, _ := authService.GetAuth()
	return username == auth.Username && verifyPassword(password, auth.Password)
}

func checkTokenAuth(c *gin.Context) bool {
	authHeader := c.GetHeader("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return false
	}
	token := splitToken[1]
	auth, _ := authService.GetAuth()
	return token == auth.PluginToken
}

func verifyPassword(password, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}
