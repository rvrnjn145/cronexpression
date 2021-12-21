package parserType

import (
	"fmt"
	"strconv"
	"strings"
)

// Returns the values seperated by comma

type CommaParser struct {
}

//check the validity of tokenized expression for  CommaParser i.e. checking 1,3
func (c *CommaParser) Check(str string) bool {
	if commaParserRegex.MatchString(str) {
		return true
	}
	return false
}

//Evaluate converts tokenized expression  to verbose for i.e.  1,3,5 -> 1 3 5
func (c *CommaParser) Evaluate(expType int, str string) (string, error) {

	low := minMaxMap[expType].min
	high := minMaxMap[expType].max
	res := ""
	parts := strings.Split(str, ",")

	for _, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			return "", fmt.Errorf("invalid CRON expression: %v", err)

		}

		if val < low || val > high {
			return "", fmt.Errorf("invalid CRON expression: %v", err)
		}
		res += ` ` + part
	}
	res = strings.Trim(res, " ")
	return res, nil
}
