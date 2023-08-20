package helpers

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken2(user interface{}, access_token string) string {
	expirationTime := time.Now().Add(1 * time.Minute).Unix()
	claims := jwt.MapClaims{
		"user":         user,
		"access_token": access_token,
		"expired_at":   expirationTime,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	singnedToken, _ := parseToken.SignedString([]byte(os.Getenv("AUTHSECRETKEY")))
	return singnedToken
	// claims := jwt.MapClaims{}

	// // Get the type of the user object
	// userType := reflect.TypeOf(user)
	// if userType.Kind() != reflect.Struct {
	// 	return "invalid user data type, expected struct"
	// }

	// // Get the value of the user object
	// userValue := reflect.ValueOf(user)

	// // Iterate through the fields of the struct
	// for i := 0; i < userType.NumField(); i++ {
	// 	field := userType.Field(i)
	// 	fieldName := field.Tag.Get("json") // Convert field name to lowercase (optional)

	// 	// Get the field value
	// 	fieldValue := userValue.Field(i)

	// 	// Add the field value to the claims map
	// 	claims[fieldName] = fieldValue.Interface()
	// }

	// // Generate the token with the claims
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// // Get the secret key from environment variable (replace "AUTHSECRETKEY" with your actual secret key environment variable)
	// secretKey := os.Getenv("AUTHSECRETKEY")
	// if secretKey == "" {
	// 	return "secret key is not set"
	// }

	// // Sign the token with the secret key and get the complete, signed token as a string
	// signedToken, err := token.SignedString([]byte(secretKey))
	// if err != nil {
	// 	return "invalid signed token"
	// }

	// return signedToken
}

func GenerateToken1(id int, email string) string {
	expirationTime := time.Now().Add(1 * time.Hour).Unix()
	claims := jwt.MapClaims{
		"user_id":    id,
		"email":      email,
		"expired_at": expirationTime,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	singnedToken, _ := parseToken.SignedString([]byte(os.Getenv("AUTHSECRETKEY")))
	return singnedToken
}

func GenerateToken(id int, email string, phone string) string {
	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	claims := jwt.MapClaims{
		"user_id":            id,
		"email":              email,
		"phone":              phone,
		"expired_at":         expirationTime,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	singnedToken, _ := parseToken.SignedString([]byte(os.Getenv("AUTHSECRETKEY")))
	return singnedToken
}

func VerifyToken(c *gin.Context) (interface{}, string, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		return nil, "", errResponse
	}
	stringToken := strings.Split(headerToken, " ")[1]
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(os.Getenv("AUTHSECRETKEY")), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, "", errResponse
	}
	return token.Claims.(jwt.MapClaims), stringToken, nil
}

func VerifyToken1(stringToken string) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(os.Getenv("AUTHSECRETKEY")), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}
	return token.Claims.(jwt.MapClaims), nil
}
