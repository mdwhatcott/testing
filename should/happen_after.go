package should

import "time"

// HappenAfter ensures that the first time value happens after the second.
func HappenAfter(actual any, expected ...any) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}
	err = validateType(actual, time.Time{})
	if err != nil {
		return err
	}
	err = validateType(expected[0], time.Time{})
	if err != nil {
		return err
	}
	AFTER := actual.(time.Time)
	BEFORE := expected[0].(time.Time)
	if AFTER.After(BEFORE) {
		return nil
	}
	return failure("unfortunately,\n%v did not happen after\n%v", AFTER, BEFORE)
}
