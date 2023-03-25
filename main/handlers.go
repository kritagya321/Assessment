package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// Maps ReceiptId to Receipt.
var mapper = make(map[string]Receipt)

// Generates ID for given receipt and stores the id -> receipt in mapper.
// Parameter 'c' is the context of current HTTP request.
// Returns    : 'id' for receipt in JSON format.
func getId(c echo.Context) (err error) {
	var requestRecepit Receipt

	if err = c.Bind(&requestRecepit); !validateReceipt(requestRecepit) || err != nil {
		return c.String(http.StatusBadRequest, "The receipt is invalid")
	}

	var receiptId = getUUID()
	mapper[receiptId] = requestRecepit
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": receiptId,
	})
}

// Calculates points for given id.
// Parameter 'c' is the context of current HTTP request.
// Returns: 'points' for receipt in JSON format
func getPoints(c echo.Context) (err error) {
	var id string = c.Param("id")
	for idex, receipt := range mapper {
		if id == idex {
			receipt = mapper[id]
			var points = calculatePoints(receipt)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"points": points,
			})
		}
	}
	return c.String(http.StatusBadRequest, "The receipt is invalid")
}
