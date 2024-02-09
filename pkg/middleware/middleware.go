package middleware

// Middleware - middleware config struct
type Middleware struct {
	SignKey []byte
}

// New - create new middleware config struct
func New(key []byte) *Middleware {
	return &Middleware{
		SignKey: key,
	}
}
