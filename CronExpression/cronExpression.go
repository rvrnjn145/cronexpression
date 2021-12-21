package cronExpression

var (
	// DefaultClient is the default client.
	DefaultClient Client
)

// ExpressionDescriptor represents the CRON expression descriptor.
type ClientImpl struct {
	logger Logger
	parser Parser
}

type Client interface {
	CronExpressionDescriptor(expr string) (desc string, err error)
}

func Init() {
	DefaultClient = NewClient()
}

func NewClient() Client {
	exprDesc := &ClientImpl{
		parser: &ParserImpl{},
		logger: &loggerImpl{LoggerFlag: true},
	}
	return exprDesc
}

func (e *ClientImpl) CronExpressionDescriptor(expr string) (string, error) {
	descriptor, err := e.parser.ParseExpression(expr)
	if err != nil {
		return "", err
	}
	e.logger.Printf("14-space rightside padded", descriptor)
	return descriptor[0], nil
}
