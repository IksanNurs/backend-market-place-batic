package handler

import (
	"fmt"
	"os"
	"time"

	"e-commerce/database"
	"e-commerce/helpers"
	model "e-commerce/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	midtrans "github.com/veritrans/go-midtrans"

	"github.com/gin-gonic/gin"
)

func GetOrderURL(orderID string, amount int, useremail string, username string) (string, string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("SERVER_KEY")
	midclient.ClientKey = os.Getenv("CLIENT_KEY")

	if os.Getenv("MIDTRANS_ENV") != "midtrans.Production" {
		midclient.APIEnvType = midtrans.Sandbox
	} else {
		midclient.APIEnvType = midtrans.Production
	}

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: useremail,
			FName: username,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", "", err
	}

	return snapTokenResp.RedirectURL, snapTokenResp.Token, nil
}

func PostOrder(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int32(userData["user_id"].(float64))
	useremail := userData["email"].(string)
	username := userData["username"].(string)
	var paymentURL string
	var paymentToken string
	var payment model.Order1
	var payment1 model.Order4
	var inputpayment model.Order6
	if err := c.ShouldBindJSON(&inputpayment); err != nil {
		response := helpers.APIResponse(err.Error(), http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println(inputpayment.Amount)

	formatUUID := "ec -" + uuid.New().String()
	payment.PackageID = inputpayment.PackageID
	payment.UserID = &userID
	payment.UUID = &formatUUID

	p := model.Order1{
		PackageID: inputpayment.PackageID,
		UserID:    &userID,
	}
	err := db.Debug().Where(p).First(&payment1).Error
	if err == nil {
		if payment1.PaidAt == nil && payment1.ExpiredAt == nil && payment1.MidtransPaymentURL != nil {
			response := helpers.APIResponse("berhasil menambah data payment", http.StatusOK, gin.H{"midtrans_payment_url": payment1.MidtransPaymentURL})
			c.JSON(http.StatusOK, response)
			return
		}
	}

	err = db.Debug().Create(&payment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paymentURL, paymentToken, err = GetOrderURL(*payment.UUID, inputpayment.Amount, useremail, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.Debug().Model(model.Order3{}).Where("id=?", payment.ID).Updates(map[string]interface{}{
		"midtrans_payment_url":   paymentURL,
		"Midtrans_payment_token": paymentToken,
		"amount":                inputpayment.Amount,
		"updated_at":             time.Now().Unix(),
		"checked_out_at":         time.Now().Unix(),
	}).Error
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusInternalServerError, errorMessage)
		return
	}

	response := helpers.APIResponse("berhasil menambah data payment", http.StatusOK, gin.H{"midtrans_payment_url": paymentURL})
	c.JSON(http.StatusOK, response)
}

func GetAllOrder(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := int32(userData["user_id"].(float64))
	var payment []model.Order
	err := db.Debug().
		Preload("Package").
		Preload("Package.Product").
		Where("user_id=?", userID).
		Order("id desc").
		Find(&payment).
		Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(payment) == 0 {
		payment = nil
	}
	response := helpers.APIResponse("berhasil menampilkan data payment", http.StatusOK, payment)
	c.JSON(http.StatusOK, response)

}

// func GetOneOrderUUID(c *gin.Context) {
// 	db := database.GetDB()
// 	userData := c.MustGet("userData").(jwt.MapClaims)
// 	userID := int32(userData["user_id"].(float64))
// 	uuid := c.Param("uuid")
// 	var payment model.Order
// 	var paymentquestion model.Order5
// 	var exam_id model.UpdateExam
// 	err := db.Debug().
// 		Preload("Package").
// 		Preload("Bundle").
// 		Where("user_id=? AND uuid=?", userID, uuid).
// 		First(&payment).
// 		Error
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
// 		return
// 	}

// 	if payment.PackageID != nil {
// 		err = db.Debug().Table("package_question").
// 			Select("COALESCE(COUNT(id), 0) as countquestion").
// 			Where("package_id=?", payment.PackageID).
// 			First(&paymentquestion).Error

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
// 			return
// 		}
// 		err = db.Table("exam").Joins("LEFT JOIN `order` ON exam.order_id=order.id").Select("exam.id").Where("exam.package_id = ? AND order.user_id=?", payment.PackageID, userID).First(&exam_id).Error
// 		if err != nil {
// 			exam_id.ID = 0
// 		}
// 	}
// 	response := helpers.APIResponse("berhasil menampilkan data payment", http.StatusOK, gin.H{"exam_id": exam_id.ID, "countquestion": paymentquestion.CountQuestion, "payment": payment})
// 	c.JSON(http.StatusOK, response)

// }
