package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"move/common"
	"move/model"
	"move/response"
	"net/http"
	"strconv"
)

func OrderList(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	limitStr := ctx.Query("limit")

	page, limit, err := validateOrderListRequest(pageStr, limitStr)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error())
		return
	}

	db := common.GetDB()
	var orders []model.Orders

	result := db.Model(model.Orders{}).Select("order_id, distance, status").Find(&orders).Where("status = ?", common.STATUS_UNASSIGNED).Order("create_time DESC").Offset((page - 1) * limit).Limit(limit)

	if result.Error != nil {
		response.Fail(ctx, http.StatusInternalServerError, result.Error.Error())
		return
	}

	var orderCnt = len(orders)
	orderList := make([]map[string]interface{}, orderCnt)
	for _, v := range orders {
		order := make(map[string]interface{})
		order["order_id"] = v.OrderId
		order["distance"] = v.Distance
		order["status"] = v.Status
		order["create_time"] = v.CreatedAt
		orderList = append(orderList, order)
	}

	ctx.JSON(http.StatusOK, orderList)
	return
}

func validateOrderListRequest(pageStr, limitStr string) (int, int, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid page value: %v", err)
	}

	if page < 1 {
		return 0, 0, err
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return 0, 0, err
	}

	if limit < 1 {
		return 0, 0, fmt.Errorf("limit must be greater than or equal to 1")
	}

	return page, limit, nil
}
