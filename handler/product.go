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

func PostProduct(c *gin.Context) {
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
	price, _ := strconv.Atoi(c.PostForm("price"))
	weight, _ := strconv.Atoi(c.PostForm("weight"))
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	category_id, _ := strconv.Atoi(c.PostForm("category_id"))
	product := models1.Product{Name: c.PostForm("name"), Image: fileName, Price: int32(price), CategoryID: int32(category_id), Weight: int32(weight), Stock: int32(stock), Deskripsi: c.PostForm("deskripsi"), UserID: int32(userID)}
	err = db.Debug().Create(&product).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menambah data product!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.APIResponse("berhasil menambah data product!", http.StatusOK, gin.H{
		"product": product,
	})
	if file != nil {
		path := fmt.Sprintf("img/product/%d-%s", product.ID, file.Filename)
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helpers.APIResponse("gagal menambah data product!", http.StatusInternalServerError, errorMessage)
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}
	c.JSON(http.StatusOK, response)
}

func PutProduct(c *gin.Context) {
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

	price, _ := strconv.Atoi(c.PostForm("price"))
	weight, _ := strconv.Atoi(c.PostForm("weight"))
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	category_id, _ := strconv.Atoi(c.PostForm("category_id"))
	product := models1.InputProduct{Name: c.PostForm("name"), Image: fileName, Price: int32(price), CategoryID: int32(category_id), Weight: int32(weight), Stock: int32(stock), Deskripsi: c.PostForm("deskripsi"), UserID: int32(userID)}
	// contentType := helpers.GetContentType(c)
	// if contentType == appJSON {
	// 	if err := c.ShouldBindJSON(&product); err != nil {
	// 		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
	// 		c.JSON(http.StatusBadRequest, response)
	// 		return
	// 	}
	// } else {
	// 	if err := c.ShouldBind(&product); err != nil {
	// 		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
	// 		c.JSON(http.StatusBadRequest, response)
	// 		return
	// 	}
	// }
	err = db.Debug().Model(&product).Where("id=? AND user_id=?", c.PostForm("id"), userID).Updates(&product).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal update product!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	err = db.Table("products").Where("id = ?", c.PostForm("id")).UpdateColumn("updated_at", time.Now().Unix()).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		return
	}
	if file != nil {
		path := fmt.Sprintf("img/product/%s-%s", c.PostForm("id"), file.Filename)
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			response := helpers.APIResponse("gagal menambah data product!", http.StatusInternalServerError, errorMessage)
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}
	response := helpers.APIResponse("berhasil update data product!", http.StatusOK, gin.H{
		"product": product,
	})
	c.JSON(http.StatusOK, response)
}

func GetProductSales(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["user_id"].(float64))
	var product []models1.Product
	err := db.Debug().
		Where("user_id=?", userID).
		Find(&product).
		Error

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menampilkan data product!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("Berhasil menampilkan data product", http.StatusOK, gin.H{
		"product": product,
	})
	c.JSON(http.StatusOK, response)
}

func GetProduct(c *gin.Context) {
	db := database.GetDB()
	var product []models1.Product
	err := db.Debug().
		Preload("Size").
		Preload("Motif").
		Find(&product).
		Error

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menampilkan data product!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("Berhasil menampilkan data product", http.StatusOK, gin.H{
		"product": product,
	})
	c.JSON(http.StatusOK, response)
}

func GetOneProduct(c *gin.Context) {
	db := database.GetDB()
	paramproduct := c.Param("id")
	var product []models1.Product
	err := db.Debug().
		Preload("Size").
		Preload("Motif").
		Where("name LIKE ?", "%"+paramproduct+"%").
		Find(&product).
		Error

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menampilkan data product!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("Berhasil menampilkan data product", http.StatusOK, gin.H{
		"product": product,
	})
	c.JSON(http.StatusOK, response)
}

func GetOneProductDetail(c *gin.Context) {
	db := database.GetDB()
	paramproduct := c.Param("id")
	var product []models1.Product
	err := db.Debug().
		Preload("Size").
		Preload("Motif").
		Where("id = ?", paramproduct).
		Find(&product).
		Error

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menampilkan data product!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("Berhasil menampilkan data product", http.StatusOK, gin.H{
		"product": product,
	})
	c.JSON(http.StatusOK, response)
}

func GetOneProductCategory(c *gin.Context) {
	db := database.GetDB()
	paramproduct := c.Param("id")
	var product []models1.Product
	err := db.Debug().
		Preload("Size").
		Preload("Motif").
		Where("category_id = ?", paramproduct).
		Find(&product).
		Error

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helpers.APIResponse("gagal menampilkan data product!", http.StatusInternalServerError, errorMessage)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helpers.APIResponse("Berhasil menampilkan data product", http.StatusOK, gin.H{
		"product": product,
	})
	c.JSON(http.StatusOK, response)
}
