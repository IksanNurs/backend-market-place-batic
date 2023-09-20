package handler

import (
	"e-commerce/database"
	"e-commerce/helpers"
	models1 "e-commerce/models"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PostCategory(c *gin.Context) {
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
	}

	category := models1.Category{Name: c.PostForm("name"), Image: fileName, UserID: int32(userID)}
	err = db.Debug().Create(&category).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menambah data category!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if file != nil {
		path := fmt.Sprintf("img/category/%d-%s", category.ID, file.Filename)
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helpers.APIResponse("gagal menambah data category!", http.StatusInternalServerError, errorMessage)
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	response := helpers.APIResponse("berhasil menambah data kategori!", http.StatusOK, gin.H{
		"category": category,
	})
	c.JSON(http.StatusOK, response)
}

func PutCategory(c *gin.Context) {
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
	}

	category := models1.Category{Name: c.PostForm("name"), Image: fileName, UserID: int32(userID)}
	// contentType := helpers.GetContentType(c)
	// if contentType == appJSON {
	// 	if err := c.ShouldBindJSON(&category); err != nil {
	// 		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
	// 		c.JSON(http.StatusBadRequest, response)
	// 		return
	// 	}
	// } else {
	// 	if err := c.ShouldBind(&category); err != nil {
	// 		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
	// 		c.JSON(http.StatusBadRequest, response)
	// 		return
	// 	}
	// }
	err = db.Debug().Model(&category).Where("id=? AND user_id=?", c.PostForm("id"), userID).Updates(&category).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal update category!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = db.Table("categories").Where("id = ?", c.PostForm("id")).UpdateColumn("updated_at", time.Now().Unix()).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		return
	}
	if file != nil {
		path := fmt.Sprintf("img/category/%s-%s", c.PostForm("id"), file.Filename)
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helpers.APIResponse("gagal menambah data category!", http.StatusInternalServerError, errorMessage)
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}
	response := helpers.APIResponse("berhasil update data kategori!", http.StatusOK, gin.H{
		"category": category,
	})
	c.JSON(http.StatusOK, response)
}

func GetCategorySales(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["user_id"].(float64))
	var category []models1.Category
	err := db.Debug().
		Where("user_id=?", userID).
		Find(&category).
		Error

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menampilkan data category!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("Berhasil menampilkan data category", http.StatusOK, gin.H{
		"category": category,
	})
	c.JSON(http.StatusOK, response)
}

func GetCategory(c *gin.Context) {
	db := database.GetDB()
	var category []models1.Category
	err := db.Debug().
		Find(&category).
		Error

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menampilkan data category!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("Berhasil menampilkan data category", http.StatusOK, gin.H{
		"category": category,
	})
	c.JSON(http.StatusOK, response)
}

func GetOneCategory(c *gin.Context) {
	db := database.GetDB()
	paramcategory := c.Param("id")
	var category []models1.Category
	err := db.Debug().
		Where("name LIKE ?", "%"+paramcategory+"%").
		Find(&category).
		Error

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menampilkan data category!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("Berhasil menampilkan data category", http.StatusOK, gin.H{
		"category": category,
	})
	c.JSON(http.StatusOK, response)
}
