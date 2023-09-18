package handler

import (
	"e-commerce/database"
	"e-commerce/helpers"
	models1 "e-commerce/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostPackage(c *gin.Context) {
	db := database.GetDB()
	var package1 models1.InputPackage
	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		if err := c.ShouldBindJSON(&package1); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	} else {
		if err := c.ShouldBind(&package1); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	err := db.Debug().Create(&package1).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menambah data package1!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("berhasil menambah data package1!", http.StatusOK, gin.H{
		"package1": package1,
	})
	c.JSON(http.StatusOK, response)
}

func PutPackage(c *gin.Context) {
	db := database.GetDB()
	var package1 models1.UpdatePackage
	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		if err := c.ShouldBindJSON(&package1); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	} else {
		if err := c.ShouldBind(&package1); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	err := db.Debug().Model(&package1).Where("id=?", package1.ID).Updates(&package1).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal update package1!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = db.Table("package").Where("id = ?", package1.ID).UpdateColumn("updated_at", time.Now().Unix()).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		return
	}
	response := helpers.APIResponse("berhasil update data package1!", http.StatusOK, gin.H{
		"package1": package1,
	})
	c.JSON(http.StatusOK, response)
}


func GetOnePackage(c *gin.Context) {
	db := database.GetDB()
	paramproduct := c.Param("id")
	var product models1.Package
	err := db.Debug().
	    Preload("Method").
	    Preload("Shipping").
	    Preload("Motif").
	    Preload("Size").
	    Preload("Product").
		Where("id = ?", paramproduct).
		First(&product).
		Error

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menampilkan data package!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("Berhasil menampilkan data package", http.StatusOK, gin.H{
		"package": product,
	})
	c.JSON(http.StatusOK, response)
}
// func GetPackage(c *gin.Context) {
// 	db := database.GetDB()
// 	var package1 []models1.Package
// 	err := db.Debug().
// 		Find(&package1).
// 		Error

// 	if err != nil {
// 		errorMessage := gin.H{"errors": err.Error()}
// 		response := helpers.APIResponse("gagal menampilkan data package1!", http.StatusInternalServerError, errorMessage)
// 		c.JSON(http.StatusInternalServerError, response)
// 		return
// 	}

// 	response := helpers.APIResponse("Berhasil menampilkan data package1", http.StatusOK, gin.H{
// 		"package1": package1,
// 	})
// 	c.JSON(http.StatusOK, response)
// }

// func GetOnePackage(c *gin.Context) {
// 	db := database.GetDB()
// 	parampackage1 := c.Param("id")
// 	var package1 []models1.Package
// 	err := db.Debug().
// 		Where("product_id = ?", parampackage1).
// 		Find(&package1).
// 		Error

// 	if err != nil {
// 		errorMessage := gin.H{"errors": err.Error()}
// 		response := helpers.APIResponse("gagal menampilkan data package1!", http.StatusInternalServerError, errorMessage)
// 		c.JSON(http.StatusInternalServerError, response)
// 		return
// 	}

// 	response := helpers.APIResponse("Berhasil menampilkan data package1", http.StatusOK, gin.H{
// 		"package1": package1,
// 	})
// 	c.JSON(http.StatusOK, response)
// }
