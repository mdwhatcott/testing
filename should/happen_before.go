package should

import "time"

// HappenBefore ensures that the first time value happens before the second.
func HappenBefore(actual any, expected ...any) error {
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
	BEFORE := actual.(time.Time)
	AFTER := expected[0].(time.Time)
	if BEFORE.Before(AFTER) {
		return nil
	}
	return failure("unfortunately,\n%v did not happen before\n%v", BEFORE, AFTER)
}
