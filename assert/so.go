package assert

type assertion func(actual interface{}, expected ...interface{}) error

func So(actual interface{}, assertion assertion, expected ...interface{}) error {
	return assertion(actual, expected...)
}
