package main

/*
File contains utility functions to generate UUID and caluclate points for receipts
*/
import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
)

// Returns UUID
func getUUID() string {
	var id = uuid.New()
	return id.String()
}

// Check if the num is a round dollar amount with no cents.
func isRoundDollarAmount(num float64) bool {
	return math.Mod(num, 1) == 0 && num > 0
}

// Calculates point based on the rules provided.
// Takes in Receipt as parameter
// Returns: Total points in int for receipt
func calculatePoints(receipt Receipt) int {
	//One point for every alphanumeric character in retailer name
	var points int = 0
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}

	//50 points if the total is a round dollar amount with no cents.
	total := stringToFloat64(receipt.Total)
	if isRoundDollarAmount(total) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt.
	points += len(receipt.Items) / 2 * 5

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
	//The result is the number of points earned.
	for _, item := range receipt.Items {
		var newString string = strings.TrimSpace(item.ShortDescription)
		var descriptionLength int = len(newString)
		if descriptionLength%3 == 0 {
			amount := stringToFloat64(item.Price)
			points += int(math.Ceil(amount * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd. Date Format: YYYY-MM-DD.
	// DateTime Layout https://pkg.go.dev/time#pkg-constants.
	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	var day int = date.Day()
	if err == nil && day%2 == 1 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	time, err := time.Parse("15:04:05", receipt.PurchaseTime+":00")
	var hour, minute int = time.Hour(), time.Minute()
	if err == nil && ((hour == 14 && minute > 0) || (hour == 15)) {
		points += 10
	}
	return points
}

// Check if requestReceipt contains null values for any fields.
func validateReceipt(requestRecepit Receipt) bool {
	if len(requestRecepit.Items) != 0 && requestRecepit.Retailer != "" && requestRecepit.PurchaseDate != "" &&
		requestRecepit.PurchaseTime != "" && requestRecepit.Total != "" {
		return true
	}
	return false
}

// Convert string literals to float amount.
// Takes amount as String and returns Float64.
func stringToFloat64(amount string) float64 {
	total, _ := strconv.ParseFloat(amount, 64)
	return total
}
