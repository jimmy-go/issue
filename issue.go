// Package issue it's a tool for micro ticket tracking.
package issue

// Issue type.
type Issue struct {
	// TODO;
}

// New returns a new Issue client.
func New(s string) (*Issue, error) {
	return &Issue{}, nil
}
