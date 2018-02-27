package issue

import (
	"context"
	"testing"

	"github.com/jimmy-go/issue/lite"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestNew(t *testing.T) {

	// Prepare storage medium.

	col, err := lite.Connect(":memory:")
	assert.Nil(t, err, "sqlite connection")
	assert.NotNil(t, col)

	// Make a new Issue tracker.

	c, err := New(col)
	assert.Nil(t, err, "new tracker")
	assert.NotNil(t, c)
}

func TestStoreAndRetrieve(t *testing.T) {

	// Prepare storage medium.

	col, err := lite.Connect(":memory:")
	assert.Nil(t, err)
	assert.NotNil(t, col)

	// Make a new Issue tracker.

	c, err := New(col)
	assert.Nil(t, err)
	assert.NotNil(t, c)

	ctx := context.TODO()

	// Add ticket issue.
	issueTitle := "Error adding user"
	issueContent := `Get some payload : {"msg":"hi error"}`

	id, err := c.Add(ctx, issueTitle, issueContent)
	assert.Nil(t, err)
	assert.EqualValues(t, true, id > 0)

	id2, err := c.Add(ctx, issueTitle, issueContent)
	assert.Nil(t, err)
	assert.EqualValues(t, true, id2 > id)

	// Retrieve ticket issue.

	title, actual, err := c.Retrieve(ctx, id)
	assert.Nil(t, err)
	assert.EqualValues(t, issueContent, actual)
	assert.EqualValues(t, issueTitle, title)
}
