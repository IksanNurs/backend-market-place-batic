package handler

import (
	"e-commerce/database"
	"e-commerce/helpers"
	models1 "e-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	appJSON = "application/json"
)

type User1 struct {
	IsSales int `json:"is_sales"`
}

func Register(c *gin.Context) {
	db := database.GetDB()

	var inputuser models1.InputUser
	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		if err := c.ShouldBindJSON(&inputuser); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	} else {
		if err := c.ShouldBind(&inputuser); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	user := models1.User{
		Name:         inputuser.Name,
		Email:        inputuser.Email,
		Phone:        inputuser.Phone,
		PasswordHash: inputuser.PasswordHash,
	}
	err := db.Debug().Create(&user).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal daftar akun!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	tokenString := helpers.GenerateToken(int(user.ID), user.Email, user.Phone)
	err = db.Table("users").Where("id = ?", user.ID).UpdateColumn("auth_key", tokenString).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal daftar akun!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("berhasil daftar akun!", http.StatusOK, gin.H{
		"token": tokenString,
		"user":  user,
	})
	c.JSON(http.StatusOK, response)
}

func Login(c *gin.Context) {
	db := database.GetDB()

	var user models1.User
	var inputuser models1.InputUser
	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		if err := c.ShouldBindJSON(&inputuser); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	} else {
		if err := c.ShouldBind(&inputuser); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	err := db.Debug().Where("email = ? OR phone = ?", inputuser.Email, inputuser.Phone).First(&user).Error
	if err != nil {
		response := helpers.APIResponse("gagal login akun!", http.StatusInternalServerError, gin.H{
			"user": User1{
				IsSales: 0,
			},
		})
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(inputuser.PasswordHash))
	if err != nil {
		//errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal login akun!", http.StatusInternalServerError, gin.H{
			"user": User1{
				IsSales: 0,
			},
		})
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("berhasil login akun!", http.StatusOK, gin.H{
		"token": user.AuthKey,
		"user":  user,
	})
	c.JSON(http.StatusOK, response)
}
