package actions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/golang-jwt/jwt"
)

func ValidateRefreshToken(refreshToken string) (string, string, error) {
	// Baca secret key untuk menandatangani token
	secret := []byte(os.Getenv("AUTHSECRETKEY"))

	// Lakukan verifikasi terhadap token yang diberikan
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return secret, nil
	})
	if err != nil {
		return "", "", err
	}

	// Ambil informasi client ID dan user ID dari token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", fmt.Errorf("invalid or expired token")
	}
	userID := claims["user_id"].(string)
	clientID := claims["client_id"].(string)

	return userID, clientID, nil
}

func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
