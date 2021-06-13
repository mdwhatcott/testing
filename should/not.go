package should

// NOT (a singleton) constrains all negated assertions to their own namespace.
var NOT negated // TODO: NOT or Not?

type negated struct{}
