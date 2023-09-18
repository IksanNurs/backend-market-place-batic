package middlewares

import (
	"e-commerce/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, _, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err":     "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		//userData := verifyToken.(jwt.MapClaims)
		// err = db.Debug().Table("user").Where("id=?", userID).First(&User).Error
		// if err != nil {
		// 	errorMessage := gin.H{"errors": err.Error()}
		// 	response := helpers.APIResponse("Unauthorized3", http.StatusUnauthorized, errorMessage)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }
		// if User.AuthKey != tokenString {
		// 	errorMessage := gin.H{"errors": "not the same"}
		// 	response := helpers.APIResponse("Unauthorized4", http.StatusUnauthorized, errorMessage)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		// jwtCreatedat := int64(userData["expired_at"].(float64))
		// timenow := time.Now().Unix()

		// if timenow > int64(jwtCreatedat) {
		// 	errorMessage := gin.H{"errors": "token expired"}
		// 	response := helpers.APIResponse("Unauthorized5", http.StatusUnauthorized, errorMessage)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		c.Set("userData", verifyToken)
		c.Next()
	}
}
