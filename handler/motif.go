package handler

import (
	"e-commerce/database"
	"e-commerce/helpers"
	models1 "e-commerce/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PostMotif(c *gin.Context) {
	db := database.GetDB()
	var fileName string
	file, err := c.FormFile("image")
	if err != nil {
		file = nil
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["user_id"].(float64))
	if file != nil {
		fileName = file.Filename
		path := fmt.Sprintf("img/motif/%d-%s", userID, file.Filename)
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helpers.APIResponse("gagal menambah data motif!", http.StatusInternalServerError, errorMessage)
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}
	price, _ := strconv.Atoi(c.PostForm("price"))
	product_id, _ := strconv.Atoi(c.PostForm("product_id"))
	motif := models1.Motif{Name: c.PostForm("name"), Image: fileName, Price: int32(price), ProductID: int32(product_id)}
	err = db.Debug().Create(&motif).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menambah data motif!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("berhasil menambah data motif!", http.StatusOK, gin.H{
		"motif": motif,
	})
	c.JSON(http.StatusOK, response)
}

func PutMotif(c *gin.Context) {
	db := database.GetDB()

	var fileName string
	file, err := c.FormFile("image")
	if err != nil {
		file = nil
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["user_id"].(float64))
	if file != nil {
		fileName = file.Filename
		path := fmt.Sprintf("img/motif/%d-%s", userID, file.Filename)
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helpers.APIResponse("gagal menambah data motif!", http.StatusInternalServerError, errorMessage)
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	price, _ := strconv.Atoi(c.PostForm("price"))
	product_id, _ := strconv.Atoi(c.PostForm("product_id"))
	motif := models1.Motif{Name: c.PostForm("name"), Image: fileName, Price: int32(price), ProductID: int32(product_id)}
	// contentType := helpers.GetContentType(c)
	// if contentType == appJSON {
	// 	if err := c.ShouldBindJSON(&motif); err != nil {
	// 		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
	// 		c.JSON(http.StatusBadRequest, response)
	// 		return
	// 	}
	// } else {
	// 	if err := c.ShouldBind(&motif); err != nil {
	// 		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
	// 		c.JSON(http.StatusBadRequest, response)
	// 		return
	// 	}
	// }
	err = db.Debug().Model(&motif).Where("id=?", c.PostForm("id")).Updates(&motif).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal update motif!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = db.Table("motifs").Where("id = ?", c.PostForm("id")).UpdateColumn("updated_at", time.Now().Unix()).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		return
	}
	response := helpers.APIResponse("berhasil update data motif!", http.StatusOK, gin.H{
		"motif": motif,
	})
	c.JSON(http.StatusOK, response)
}

func GetMotif(c *gin.Context) {
	db := database.GetDB()
	var motif []models1.Motif
	err := db.Debug().
		Find(&motif).
		Error

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menampilkan data motif!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("Berhasil menampilkan data motif", http.StatusOK, gin.H{
		"motif": motif,
	})
	c.JSON(http.StatusOK, response)
}

func GetOneMotif(c *gin.Context) {
	db := database.GetDB()
	parammotif := c.Param("id")
	var motif []models1.Motif
	err := db.Debug().
		Where("product_id = ?", parammotif).
		Find(&motif).
		Error

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menampilkan data motif!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("Berhasil menampilkan data motif", http.StatusOK, gin.H{
		"motif": motif,
	})
	c.JSON(http.StatusOK, response)
}
