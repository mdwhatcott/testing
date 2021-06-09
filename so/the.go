package so

// The is used as in so.The(thing, should.Equal, otherThing)
func The(actual interface{}, assertion assertion, expected ...interface{}) error {
	return assertion(actual, expected...)
}

type assertion func(actual interface{}, expected ...interface{}) error
