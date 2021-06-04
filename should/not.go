package should

// NOT (a singleton) constrains all negated assertions to their own namespace.
var NOT not

type not struct{}
