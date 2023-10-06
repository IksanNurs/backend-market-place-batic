package handler

import (
	"e-commerce/database"
	"e-commerce/helpers"
	model "e-commerce/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func PostNotification(c *gin.Context) {
	db := database.GetDB()
	var inputnotification model.InputMidtransNotification
	var order model.Order1
	body, _ := ioutil.ReadAll(c.Request.Body)
	contentType := helpers.GetContentType(c)
	if contentType == appJSON {
		if err := json.Unmarshal(body, &inputnotification); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	} else {
		if err := json.Unmarshal(body, &inputnotification); err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}
	err := db.Debug().Model(model.Order4{}).Where("uuid=?", inputnotification.OrderUUID).Updates(map[string]interface{}{
		"checked_out_at": time.Now().Unix(),
		"updated_at":     time.Now().Unix(),
	}).Error
	if err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = db.Debug().Where("uuid=?", inputnotification.OrderUUID).First(&order).Error
	if err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	raw := string(body)
	grossamount, _ := strconv.ParseInt(strings.Split(*inputnotification.GrossAmount, ".")[0], 10, 64)
	notification := model.MidtransNotification{
		OrderID:           &order.ID,
		OrderUUID:         inputnotification.OrderUUID,
		SignatureKey:      inputnotification.SignatureKey,
		TransactionID:     inputnotification.TransactionID,
		TransactionTime:   inputnotification.TransactionTime,
		TransactionStatus: inputnotification.TransactionStatus,
		FraudStatus:       inputnotification.FraudStatus,
		PaymentType:       inputnotification.PaymentType,
		GrossAmount:       &grossamount,
		Raw:               &raw,
	}

	err = db.Debug().Create(&notification).Error
	if err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if *notification.PaymentType == "credit_card" && *notification.TransactionStatus == "capture" && *notification.FraudStatus == "accept" {
		err = db.Debug().Model(model.Order4{}).Where("uuid=?", notification.OrderUUID).Updates(map[string]interface{}{
			"paid_at":    time.Now().Unix(),
			"updated_at": time.Now().Unix(),
		}).Error
		if err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	} else if *notification.TransactionStatus == "settlement" {
		err = db.Debug().Model(model.Order4{}).Where("uuid=?", notification.OrderUUID).Updates(map[string]interface{}{
			"paid_at":    time.Now().Unix(),
			"updated_at": time.Now().Unix(),
		}).Error
		if err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	} else if *notification.TransactionStatus == "deny" || *notification.TransactionStatus == "expire" {
		err = db.Debug().Model(model.Order4{}).Where("uuid=?", notification.OrderUUID).Updates(map[string]interface{}{
			"expired_at": time.Now().Unix(),
			"updated_at": time.Now().Unix(),
		}).Error
		if err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	} else if *notification.TransactionStatus == "cancel" {
		err = db.Debug().Model(model.Order4{}).Where("uuid=?", notification.OrderUUID).Updates(map[string]interface{}{
			"cancelled_at": time.Now().Unix(),
			"updated_at":   time.Now().Unix(),
		}).Error
		if err != nil {
			response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

	}
	response := helpers.APIResponse("berhasil menambah data notification", http.StatusOK, notification)
	c.JSON(http.StatusOK, response)
}
