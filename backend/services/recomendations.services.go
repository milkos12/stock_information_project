package services

import (
	"regexp"
	"stock_information_project/models"
	"strconv"
)

func parseFloat(value string) float64 {
	if value == "" {
		return 0
	}

	re := regexp.MustCompile(`[^0-9.]`)
	cleaned := re.ReplaceAllString(value, "")

	result, err := strconv.ParseFloat(cleaned, 64)
	if err != nil {
		return 0
	}

	return result
}

var positiveRatings = map[string]bool{
	"Strong-Buy":        true,
	"Buy":               true,
	"Overweight":        true,
	"Outperform":        true,
	"Market Outperform": true,
	"Sector Outperform": true,
	"Speculative Buy":   true,
	"Outperformer":      true,
	"Positive":          true,
}

func IsRecommended(stock models.Stock) bool {
	targetFrom := parseFloat(stock.TargetFrom)
	targetTo := parseFloat(stock.TargetTo)

	targetIncreased := targetTo > targetFrom

	var highPotential bool

	potentialUpside := ((targetTo - targetFrom) / targetFrom) * 100
	highPotential = potentialUpside > 0

	ratingWasPositive := positiveRatings[stock.RatingFrom]
	ratingIsPositive := positiveRatings[stock.RatingTo]
	ratingStillPositive := ratingWasPositive && ratingIsPositive
	ratingImproved := !ratingWasPositive && ratingIsPositive

	score := 0
	if targetIncreased {
		score++
	}
	if highPotential {
		score++
	}
	if ratingImproved || ratingStillPositive {
		score++

	}

	return score >= 2
}
