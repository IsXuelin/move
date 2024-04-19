package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"googlemaps.github.io/maps"
	"log"
	"move/common"
	"move/model"
	"move/response"
	"net/http"
)

type PlaceOrderRequest struct {
	Origin      []string `json:"origin"`
	Destination []string `json:"destination"`
}

func PlaceOrder(ctx *gin.Context) {
	var placeOrderRequest PlaceOrderRequest
	if err := ctx.Bind(&placeOrderRequest); err != nil {
		log.Println("Invalid Parameters, ", placeOrderRequest)
		response.Fail(ctx, http.StatusBadRequest, "Invalid Parameters")
		return
	}

	if !isValidCoordinates(ctx, placeOrderRequest) {
		return
	}

	orderId := uuid.New().String()

	var client *maps.Client
	var err error
	// input a real API_KEY
	client, err = maps.NewClient(maps.WithAPIKey(common.API_Key))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	r := &maps.DistanceMatrixRequest{
		Origins:      placeOrderRequest.Origin,
		Destinations: placeOrderRequest.Destination,
	}

	resp, err := client.DistanceMatrix(ctx, r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	var distance = resp.Rows[0].Elements[0].Distance.Meters

	db := common.GetDB()
	newOrder := model.Orders{
		OrderId:        orderId,
		StartLatitude:  placeOrderRequest.Origin[0],
		StartLongitude: placeOrderRequest.Origin[1],
		EndLatitude:    placeOrderRequest.Destination[0],
		EndLongitude:   placeOrderRequest.Destination[1],
		Distance:       distance,
		Status:         common.STATUS_UNASSIGNED,
		//TODO CustomerID: Get it from token
	}
	result := db.Create(newOrder)

	if result.Error != nil {
		response.Fail(ctx, http.StatusInternalServerError, result.Error.Error())
		return
	}
	response.Success(ctx, gin.H{
		"id":       orderId,
		"distance": distance,
		"status":   "UNASSIGNED",
	})
	return
}

func isValidLatitude(latStr string) bool {
	lat, err := decimal.NewFromString(latStr)
	if err != nil {
		return false
	}
	minLat, _ := decimal.NewFromString("-90")
	maxLat, _ := decimal.NewFromString("90")
	return lat.GreaterThanOrEqual(minLat) && lat.LessThanOrEqual(maxLat)
}

// isValidLongitude checks if the longitude value is within the valid range (-180 to 180 degrees).
func isValidLongitude(lonStr string) bool {
	lon, err := decimal.NewFromString(lonStr)
	if err != nil {
		return false
	}
	minLon, _ := decimal.NewFromString("-180")
	maxLon, _ := decimal.NewFromString("180")
	return lon.GreaterThanOrEqual(minLon) && lon.LessThanOrEqual(maxLon)
}

// isValidCoordinates checks if the latitude and longitude values are valid.
func isValidCoordinates(ctx *gin.Context, placeOrderRequest PlaceOrderRequest) bool {
	origin := placeOrderRequest.Origin
	destination := placeOrderRequest.Destination

	if len(origin) != 2 || len(destination) != 2 {
		response.Fail(ctx, http.StatusBadRequest, "Invalid count of origin or destination")
		return false
	}
	startLatitude := origin[0]
	startLongitude := origin[1]
	endLatitude := destination[0]
	endLongitude := destination[1]
	if !isValidLatitude(startLatitude) ||
		!isValidLongitude(startLongitude) ||
		!isValidLatitude(endLatitude) ||
		!isValidLongitude(endLongitude) {

		response.Fail(ctx, http.StatusBadRequest, "Invalid Latitude or Longitude")
		return false
	}
	return true
}
