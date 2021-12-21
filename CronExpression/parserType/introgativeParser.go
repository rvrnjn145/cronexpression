package parserType

import "fmt"

// Returns all the possible values for the allowed range

type IntrogativeParser struct {
}

//check the validity of tokenized expression for  CommaParser i.e. checking *
func (i *IntrogativeParser) Check(str string) bool {
	if str == "?" {
		return true
	}
	return false
}

//Evaluate converts tokenized expression  to verbose for all possible range i.e say for Day Of Week  * ->  1 2 3 4 5 6 7
func (i *IntrogativeParser) Evaluate(expType int, str string) (string, error) {

	if expType == int(DayOfMonth) || expType == int(DayOfWeek) {
		return "", nil
	}
	return "", fmt.Errorf("invalid CRON expression")
}
