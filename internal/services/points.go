package services

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"../model"
)

func CalculatePoints(receipt model.Receipt) int {
	points := 0

	// Rule 1: Alphanumeric characters in retailer name
	alnumCount := len(regexp.MustCompile(`[A-Za-z0-9]`).FindAllString(receipt.Retailer, -1))
	points += alnumCount

	// Rule 2: Total is a round dollar amount
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
	}

	// Rule 3: Total is a multiple of 0.25
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points based on item description length
	for _, item := range receipt.Items {
		descLength := len(strings.TrimSpace(item.ShortDescription))
		if descLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: Day in purchase date is odd
	day, _ := strconv.Atoi(strings.Split(receipt.PurchaseDate, "-")[2])
	if day%2 != 0 {
		points += 6
	}

	// Rule 7: Time between 2:00pm and 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 || (purchaseTime.Hour() == 15 && purchaseTime.Minute() < 60) {
		points += 10
	}

	return points
}
