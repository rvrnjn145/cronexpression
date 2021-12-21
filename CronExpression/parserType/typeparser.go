package parserType

import (
	"regexp"
	"strconv"
	"strings"
)

type limit struct {
	min int
	max int
}
type ExpressionType int

const (
	Minute     ExpressionType = 0
	Hour       ExpressionType = 1
	DayOfMonth ExpressionType = 2
	Month      ExpressionType = 3
	DayOfWeek  ExpressionType = 4
	Command    ExpressionType = 5
)

var (
	incrementRegex    = regexp.MustCompile(`^\*|[0-9]+\/[0-9]+$`)
	commaParserRegex  = regexp.MustCompile(`^[0-9]+(,[0-9]+)*$`)
	rangesParserRegex = regexp.MustCompile(`^[0-9]+-[0-9]+$`)
	//starParserRegex   = regexp.MustCompile("/*")

	minMaxMap = map[int]limit{
		int(Minute):     {0, 59},
		int(Hour):       {0, 23},
		int(DayOfMonth): {1, 31},
		int(Month):      {1, 12},
		int(DayOfWeek):  {1, 7},
	}
)

//TypeParser does the actual parsing and validation
type TypeParser interface {
	Check(str string) bool
	Evaluate(expType int, str string) (string, error)
}

func createRange(start, end int) string {
	res := ""
	for i := start; i <= end; i++ {
		res += ` ` + strconv.Itoa(i)
	}
	res = strings.Trim(res, " ")
	return res
}
