package handler

import (
	"e-commerce/database"
	"e-commerce/helpers"
	models1 "e-commerce/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostSize(c *gin.Context) {
	db := database.GetDB()
	var size models1.Size
	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		if err := c.ShouldBindJSON(&size); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	} else {
		if err := c.ShouldBind(&size); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	err := db.Debug().Create(&size).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menambah data size!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("berhasil menambah data size!", http.StatusOK, gin.H{
		"size": size,
	})
	c.JSON(http.StatusOK, response)
}

func PutSize(c *gin.Context) {
	db := database.GetDB()
	var size models1.Size
	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		if err := c.ShouldBindJSON(&size); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	} else {
		if err := c.ShouldBind(&size); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	err := db.Debug().Model(&size).Where("id=?", size.ID).Updates(&size).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal update size!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = db.Table("sizes").Where("id = ?", size.ID).UpdateColumn("updated_at", time.Now().Unix()).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		return
	}
	response := helpers.APIResponse("berhasil update data size!", http.StatusOK, gin.H{
		"size": size,
	})
	c.JSON(http.StatusOK, response)
}

// func GetSize(c *gin.Context) {
// 	db := database.GetDB()
// 	var size []models1.Size
// 	err := db.Debug().
// 		Find(&size).
// 		Error

// 	if err != nil {
// 		errorMessage := gin.H{"errors": err.Error()}
// 		response := helpers.APIResponse("gagal menampilkan data size!", http.StatusInternalServerError, errorMessage)
// 		c.JSON(http.StatusInternalServerError, response)
// 		return
// 	}

// 	response := helpers.APIResponse("Berhasil menampilkan data size", http.StatusOK, gin.H{
// 		"size": size,
// 	})
// 	c.JSON(http.StatusOK, response)
// }

// func GetOneSize(c *gin.Context) {
// 	db := database.GetDB()
// 	paramsize := c.Param("id")
// 	var size []models1.Size
// 	err := db.Debug().
// 		Where("product_id = ?", paramsize).
// 		Find(&size).
// 		Error

// 	if err != nil {
// 		errorMessage := gin.H{"errors": err.Error()}
// 		response := helpers.APIResponse("gagal menampilkan data size!", http.StatusInternalServerError, errorMessage)
// 		c.JSON(http.StatusInternalServerError, response)
// 		return
// 	}

// 	response := helpers.APIResponse("Berhasil menampilkan data size", http.StatusOK, gin.H{
// 		"size": size,
// 	})
// 	c.JSON(http.StatusOK, response)
// }
