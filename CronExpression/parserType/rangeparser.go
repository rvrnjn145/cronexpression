package parserType

import (
	"fmt"
	"strconv"
	"strings"
)

// Returns all the values that are allowed between the dash
type RangesParser struct {
}

//check the validity of tokenized expression for  RangesParser ie checking 1-3
func (r *RangesParser) Check(str string) bool {
	if rangesParserRegex.MatchString(str) {
		return true
	}
	return false
}

//Evaluate converts tokenized expression  to verbose for i.e.  1-3 -> 1,2,3
func (r *RangesParser) Evaluate(expType int, str string) (string, error) {
	low := minMaxMap[expType].min
	high := minMaxMap[expType].max
	parts := strings.Split(str, "-")
	start := parts[0]
	end := parts[1]

	if end == "" || start == "" {
		return "", fmt.Errorf("invalid CRON expression")
	}
	startVal, err := strconv.Atoi(start)
	if err != nil {
		return "", fmt.Errorf("invalid CRON expression: %v", err)
	}

	endValue, err := strconv.Atoi(end)
	if err != nil {
		return "", fmt.Errorf("invalid CRON expression: %v", err)
	}

	// Some error checking
	if startVal > endValue {
		return "", fmt.Errorf("invalid CRON expression")
	}
	if startVal < low || endValue > high {
		return "", fmt.Errorf("invalid CRON expression")
	}

	return createRange(startVal, endValue), nil
}
