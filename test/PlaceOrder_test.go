package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"move/controller"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPlaceOrderHappyCase(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	requestBody := `{"origin": ["55.930385", "-3.118425"], "destination": ["50.087692", "14.421150"]}`
	ctx.Request = httptest.NewRequest(http.MethodPost, "/orders/placeorder", strings.NewReader(requestBody))
	ctx.Request.Header.Set("Content-Type", "application/json")

	controller.PlaceOrder(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}

func TestPlaceOrderInvalidCount(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	requestBody := `{"origin": ["1.239945", "5.674589", "123"], "destination": ["2.342356", "6.781290"]}`
	ctx.Request = httptest.NewRequest(http.MethodPost, "/order", strings.NewReader(requestBody))
	ctx.Request.Header.Set("Content-Type", "application/json")

	controller.PlaceOrder(ctx)

	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

func TestPlaceOrderInvalidLonLat(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	requestBody := `{"origin": ["1.239945", "360"], "destination": ["2.342356", "6.781290"]}`
	ctx.Request = httptest.NewRequest(http.MethodPost, "/order", strings.NewReader(requestBody))
	ctx.Request.Header.Set("Content-Type", "application/json")

	controller.PlaceOrder(ctx)

	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}
