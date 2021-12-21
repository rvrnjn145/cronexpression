package cronExpression

import (
	"./parserType"
	"errors"
	"fmt"
	"strings"
)

var (
	InvalidExprError     = errors.New("invalid expression")
	sunday           Day = "1"
	monday           Day = "2"
	tuesday          Day = "3"
	wednesday        Day = "4"
	thursday         Day = "5"
	friday           Day = "6"
	saturday         Day = "7"

	weekdayMap = map[string]string{
		string(sunday):    "sunday",
		string(monday):    "monday",
		string(tuesday):   "tuesday",
		string(wednesday): "wednesday",
		string(thursday):  "thursday",
		string(friday):    "friday",
		string(saturday):  "saturday",
	}
)

type ParserImpl struct {
}
type Day string
type ParserEnum int

const (
	increment ParserEnum = iota
	comma
	ranges
	asterisk
	introgative
)

//factory for parserTypeObject
func (p *ParserImpl) getParserObj(r ParserEnum) parserType.TypeParser {
	switch r {
	case increment:
		return &parserType.IncrementParser{}
	case comma:
		return &parserType.CommaParser{}
	case ranges:
		return &parserType.RangesParser{}
	case asterisk:
		return &parserType.AsteriskParser{}
	case introgative:
		return &parserType.IntrogativeParser{}
	}
	return nil
}

// Parser represents the cronExpression parser.

type Parser interface {
	ParseExpression(expr string) (exprParts []string, err error)
}

//ParseExpression  completely parse the expression and  expands each field
//to show the times at which it will run
func (p *ParserImpl) ParseExpression(expr string) (exprParts []string, err error) {
	exprParts, err = p.tokenizeExpression(expr)
	if err != nil {
		return nil, fmt.Errorf("failed to extract expression parts: %v", err)
	}

	for i, expr := range exprParts {
		if i == int(parserType.Command) {
			continue
		}
		expr, err := p.parse(i, expr)

		if err != nil {
			return nil, fmt.Errorf("failed to extract expression parts: %v", err)
		}
		exprParts[i] = expr
	}
	return exprParts, nil
}

//try parsing a tokenized expression with differ Types of parser
func (p *ParserImpl) parse(i int, expr string) (exprParts string, e error) {
	for parser := 0; parser < 5; parser++ {
		obj := p.getParserObj(ParserEnum(parser))
		if obj.Check(expr) {
			return obj.Evaluate(i, expr)
		}
	}
	return expr, nil
}

//convert the expression string to tokens splits
func (p *ParserImpl) tokenizeExpression(expr string) (exprParts []string, err error) {
	if strings.TrimSpace(expr) == "" {
		return nil, InvalidExprError
	}

	expr = strings.ToLower(expr)
	exprParts = make([]string, 6, 6)
	parts := strings.Fields(expr)
	command := ""
	for i, parts := range parts {
		if i > 4 {
			command += parts
			command += " "
		}

	}
	parts[5] = strings.Trim(command, " ")

	switch {
	//case len(parts)  6:
	//	return nil, fmt.Errorf("error %v", InvalidExprError)
	case len(parts) > 6:
		copy(exprParts, append(parts, ""))

	}

	return exprParts, nil
}

func (p *ParserImpl) getDayAugmentedValue(expr string) string {
	days := strings.Split(expr, " ")
	res := ""
	for _, day := range days {
		res += ` ` + weekdayMap[day]
	}
	res = strings.Trim(res, " ")
	return res
}
