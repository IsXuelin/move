package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"move/controller"
	_ "move/controller"
	"net/http"
	"net/http/httptest"
	_ "strings"
	"testing"
)

func TestOrderList(t *testing.T) {
	//db, mock, err := sqlmock.New()
	//if err != nil {
	//	t.Fatalf("Failed to create mock DB: %v", err)
	//}
	//defer db.Close()
	//common.SetDB(db)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	ctx.Request = httptest.NewRequest(http.MethodGet, "/orders/orderlist?page=1&limit=10", nil)

	controller.OrderList(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}
