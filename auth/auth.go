package auth

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString(jwtSecret)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&credentials)

	// Validate against environment credentials
	if credentials.Username == os.Getenv("USER_NAME") && credentials.Password == os.Getenv("PASSWORD") {
		token, _ := GenerateToken(credentials.Username)
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
