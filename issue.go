// Package issue it's a tool for micro ticket tracking.
package issue

import "context"

// Collector interface gives free of implementation with any database. Examples for
// PostgreSQL and SQLite are provided as a guide.
type Collector interface {
	// Get retrieves ticket by ID.
	Get(ctx context.Context, id int64) (title, actual string, err error)

	// Save stores ticket expected and actual info and returns a new id.
	Save(ctx context.Context, titleOrExpected, actual string) (id int64, err error)

	// Append adds body content to issue.
	Append(ctx context.Context, id int64, body string) error
}

// Tracker type.
type Tracker struct {
	store Collector
}

// New returns a new Tracker client.
func New(col Collector) (*Tracker, error) {
	x := &Tracker{
		store: col,
	}

	return x, nil
}

// Add calls inner Collector method Get but has another name to prevent accidental
// refs.
func (xt *Tracker) Add(ctx context.Context, titleOrExpected, actual string) (int64, error) {
	id, err := xt.store.Save(ctx, titleOrExpected, actual)
	return id, err
}

// Retrieve calls inner Collector method Save but has another name to prevent accidental
// refs.
func (xt *Tracker) Retrieve(ctx context.Context, id int64) (string, string, error) {
	title, actual, err := xt.store.Get(ctx, id)
	return title, actual, err
}

// NewID returns a new empty issue id.
func (xt *Tracker) NewID(ctx context.Context, titleOrExpected string) (int64, error) {
	return xt.Add(ctx, titleOrExpected, "")
}

// Append append body info to issue.
func (xt *Tracker) Append(ctx context.Context, id int64, body string) error {
	return xt.store.Append(ctx, id, body)
}
