package handler

import (
	"e-commerce/database"
	"e-commerce/helpers"
	models1 "e-commerce/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostShipping(c *gin.Context) {
	db := database.GetDB()
	var shipping models1.Shipping
	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		if err := c.ShouldBindJSON(&shipping); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	} else {
		if err := c.ShouldBind(&shipping); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	err := db.Debug().Create(&shipping).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menambah data shipping!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("berhasil menambah data shipping!", http.StatusOK, gin.H{
		"shipping": shipping,
	})
	c.JSON(http.StatusOK, response)
}

func PutShipping(c *gin.Context) {
	db := database.GetDB()
	var shipping models1.Shipping
	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		if err := c.ShouldBindJSON(&shipping); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	} else {
		if err := c.ShouldBind(&shipping); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	err := db.Debug().Model(&shipping).Where("id=?", shipping.ID).Updates(&shipping).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal update shipping!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = db.Table("shippings").Where("id = ?", shipping.ID).UpdateColumn("updated_at", time.Now().Unix()).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		return
	}
	response := helpers.APIResponse("berhasil update data shipping!", http.StatusOK, gin.H{
		"shipping": shipping,
	})
	c.JSON(http.StatusOK, response)
}

// func GetShipping(c *gin.Context) {
// 	db := database.GetDB()
// 	var shipping []models1.Shipping
// 	err := db.Debug().
// 		Find(&shipping).
// 		Error

// 	if err != nil {
// 		errorMessage := gin.H{"errors": err.Error()}
// 		response := helpers.APIResponse("gagal menampilkan data shipping!", http.StatusInternalServerError, errorMessage)
// 		c.JSON(http.StatusInternalServerError, response)
// 		return
// 	}

// 	response := helpers.APIResponse("Berhasil menampilkan data shipping", http.StatusOK, gin.H{
// 		"shipping": shipping,
// 	})
// 	c.JSON(http.StatusOK, response)
// }

// func GetOneShipping(c *gin.Context) {
// 	db := database.GetDB()
// 	paramshipping := c.Param("id")
// 	var shipping []models1.Shipping
// 	err := db.Debug().
// 		Where("product_id = ?", paramshipping).
// 		Find(&shipping).
// 		Error

// 	if err != nil {
// 		errorMessage := gin.H{"errors": err.Error()}
// 		response := helpers.APIResponse("gagal menampilkan data shipping!", http.StatusInternalServerError, errorMessage)
// 		c.JSON(http.StatusInternalServerError, response)
// 		return
// 	}

// 	response := helpers.APIResponse("Berhasil menampilkan data shipping", http.StatusOK, gin.H{
// 		"shipping": shipping,
// 	})
// 	c.JSON(http.StatusOK, response)
// }
