package parserType

// Returns all the possible values for the allowed range

type AsteriskParser struct {
}

//check the validity of tokenized expression for  CommaParser i.e. checking *
func (i *AsteriskParser) Check(str string) bool {
	if str == "*" {
		return true
	}
	return false
}

//Evaluate converts tokenized expression  to verbose for all possible range i.e say for Day Of Week  * ->  1 2 3 4 5 6 7
func (i *AsteriskParser) Evaluate(expType int, str string) (string, error) {
	low := minMaxMap[expType].min
	high := minMaxMap[expType].max

	return createRange(low, high), nil
}
