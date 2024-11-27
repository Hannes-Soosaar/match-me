package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

var jwtSecret = os.Getenv("JWT_SECRET")

// Function to extract the user ID from the JWT token
func ExtractUserIDFromToken(tokenString string) (int, error) {
    // Parse the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Validate the token signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
        }
        return jwtSecret, nil // Return the secret key for verification
    })

    if err != nil || !token.Valid {
        return 0, fmt.Errorf("invalid token: %v", err)
    }

    // Extract the claims
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return 0, fmt.Errorf("unable to parse claims")
    }

    // Extract the userID from claims
    userID, ok := claims["userID"].(float64)
    if !ok {
        return 0, fmt.Errorf("userID not found in claims")
    }

    return int(userID), nil // Return the user ID as an integer
}
