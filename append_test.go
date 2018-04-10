package issue

import (
	"context"
	"testing"

	"github.com/jimmy-go/issue/lite"
	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	// Prepare storage medium.
	col, err := lite.Connect(":memory:")
	assert.Nil(t, err)
	assert.NotNil(t, col)

	// Make a new Issue tracker.
	c, err := New(col)
	assert.Nil(t, err)
	assert.NotNil(t, c)

	ctx := context.TODO()
	// Create a new id.
	id, err := c.NewID(ctx, "Normal")
	assert.Nil(t, err)
	assert.EqualValues(t, true, id > 0)

	// Append data.
	err = c.Append(ctx, id, "Hi doggy")
	assert.Nil(t, err)

	err = c.Append(ctx, id, "What a story mark")
	assert.Nil(t, err)

	// Add ticket issue.
	issueTitle := "Normal"
	issueContent := `Hi doggy
What a story mark
`

	// Retrieve ticket issue.

	title, actual, err := c.Retrieve(ctx, id)
	assert.Nil(t, err)
	assert.EqualValues(t, issueContent, actual)
	assert.EqualValues(t, issueTitle, title)
}
