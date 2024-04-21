package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"move/common"
	"move/model"
	"move/response"
	"net/http"
)

type TakeOrderRequest struct {
	Status string `json:"status"`
}

func TakeOrder(ctx *gin.Context) {
	var takeOrderRequest TakeOrderRequest

	orderId := ctx.Param("id")

	if err := ctx.Bind(&takeOrderRequest); err != nil {
		log.Println("Invalid Parameters, ", takeOrderRequest)
		response.Fail(ctx, http.StatusBadRequest, "Invalid Parameters")
		return
	}
	status, err := validateTakeOrderRequest(ctx, takeOrderRequest.Status)
	if err == false {
		return
	}
	db := common.GetDB()

	var orders []model.Orders
	result1 := db.Select("order_id, status").Where("order_id = ? and status = ?", orderId, common.STATUS_UNASSIGNED).Find(&orders)
	if result1.Error != nil {
		response.Fail(ctx, http.StatusInternalServerError, result1.Error.Error())
		return
	}
	if len(orders) == 0 {
		response.Fail(ctx, http.StatusBadRequest, "Wrong order info")
		return
	}
	alterField := make(map[string]interface{})
	alterField["status"] = status

	result := db.Model(&model.Orders{}).Where("order_id = ? and status = ?", orderId, common.STATUS_UNASSIGNED).Update(alterField)
	//TODO update DriverID
	if result.Error != nil {
		response.Fail(ctx, http.StatusInternalServerError, result.Error.Error())
		return
	}
	if result.RowsAffected == 1 {
		response.Success(ctx, gin.H{
			"status": "SUCCESS",
		})
		return
	}

	response.Fail(ctx, http.StatusOK, "The order has been taken by others")

	return
}

func validateTakeOrderRequest(ctx *gin.Context, status string) (int, bool) {
	if status != "TAKEN" {
		response.Fail(ctx, http.StatusBadRequest, "Invalid Parameters: status")
		return 0, false
	}
	return 1, true
}
