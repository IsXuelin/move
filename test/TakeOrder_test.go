package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"move/controller"
)

func TestHappyCase(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	requestBody := `{"status": "TAKEN"}`
	ctx.Request = httptest.NewRequest(http.MethodPatch, "/orders/111111", strings.NewReader(requestBody))
	ctx.Request.Header.Set("Content-Type", "application/json")

	controller.TakeOrder(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}

func TestInvalidStatus(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	requestBody := `{"status": "Invalid"}`
	ctx.Request = httptest.NewRequest(http.MethodPatch, "/order/123", strings.NewReader(requestBody))
	ctx.Request.Header.Set("Content-Type", "application/json")

	controller.TakeOrder(ctx)

	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

func TestInvalidOrderId(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	requestBody := `{"status": "Invalid"}`
	ctx.Request = httptest.NewRequest(http.MethodPatch, "/order", strings.NewReader(requestBody))
	ctx.Request.Header.Set("Content-Type", "application/json")

	controller.TakeOrder(ctx)

	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}
