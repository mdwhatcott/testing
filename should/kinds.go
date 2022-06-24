package should

import "reflect"

var integerKinds = []reflect.Kind{
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
}

var numericKinds = map[reflect.Kind]struct{}{
	reflect.Int:     {},
	reflect.Int8:    {},
	reflect.Int16:   {},
	reflect.Int32:   {},
	reflect.Int64:   {},
	reflect.Uint:    {},
	reflect.Uint8:   {},
	reflect.Uint16:  {},
	reflect.Uint32:  {},
	reflect.Uint64:  {},
	reflect.Float32: {},
	reflect.Float64: {},
}

func isNumeric(v any) bool {
	_, found := numericKinds[reflect.TypeOf(v).Kind()]
	return found
}

var kindsWithLength = []reflect.Kind{
	reflect.Map,
	reflect.Chan,
	reflect.Array,
	reflect.Slice,
	reflect.String,
}

var containerKinds = []reflect.Kind{
	reflect.Map,
	reflect.Array,
	reflect.Slice,
	reflect.String,
}

var orderedContainerKinds = []reflect.Kind{
	reflect.Array,
	reflect.Slice,
	reflect.String,
}
