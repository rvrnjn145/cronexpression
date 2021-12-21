package parserType

import (
	"fmt"
	"strconv"
	"strings"
)

// Returns incremented values from the defined and start

type IncrementParser struct {
}

//check the validity of tokenized expression for  IncrementParser i.e. checking 1/3
func (i *IncrementParser) Check(str string) bool {
	if incrementRegex.MatchString(str) && len(str) > 1 {
		return true
	}
	return false
}

//Evaluate converts tokenized expression  to verbose for i.e.  1/3 -> 1 4 7 10 13 17 20 23..
func (i *IncrementParser) Evaluate(expType int, str string) (string, error) {
	low := minMaxMap[expType].min
	high := minMaxMap[expType].max
	res := ""
	startingVal := 0
	incrementingVal := 1
	parts := strings.Split(str, "/")
	start := parts[0]
	increment := parts[1]
	flaghyphen := false
	for _, s := range start {
		if s == '-' { // 3-6
			flaghyphen = true
		}
	}
	if flaghyphen {
		highLowrange := strings.Split(start, "-")
		startVal, err := strconv.Atoi(highLowrange[0])
		if err != nil {
			return "", fmt.Errorf("invalid CRON expression: %v", err)
		}
		startingVal = startVal
		endVal, err := strconv.Atoi(highLowrange[1])
		if err != nil {
			return "", fmt.Errorf("invalid CRON expression: %v", err)
		}
		high = endVal

	} else {

		if start != "*" {
			startVal, err := strconv.Atoi(start)
			if err != nil {
				return "", fmt.Errorf("invalid CRON expression: %v", err)
			}
			startingVal = startVal
		} else {
			startingVal = low
		}
	}

	incrementVal, err := strconv.Atoi(increment)
	if err != nil {
		return "", fmt.Errorf("invalid CRON expression: %v", err)
	}
	incrementingVal = incrementVal
	if startingVal < low || startingVal > high {
		return "", fmt.Errorf("invalid CRON expression")
	}
	for i := startingVal; i <= high; i += incrementingVal {
		part := strconv.Itoa(i)
		res += ` ` + part
	}
	res = strings.Trim(res, " ")
	return res, nil
}
