// Package lite contains sqlite sample adapter.
package lite

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// Lite type implements issue.Collector
type Lite struct {
	db *sqlx.DB
}

// Connect returns a new SQLite adapter that implements issue.Collector
func Connect(connectionURL string) (*Lite, error) {
	db, err := sqlx.Connect("sqlite3", connectionURL)
	if err != nil {
		return nil, err
	}
	l := &Lite{db}
	if _, err := getVersion(l.db); err != nil {
		// Schema doesn't exists generate tables.
		if er := migrateSchema(l.db); er != nil {
			return nil, er
		}
	}
	return l, nil
}

// Get implements Collector.
func (li *Lite) Get(ctx context.Context, id int64) (string, string, error) {
	var m Item
	err := li.db.GetContext(ctx, &m, `
		SELECT actual,expected FROM issues WHERE id=?
	`, id)
	return m.Expected, m.Actual, err
}

// Item type.
type Item struct {
	Actual   string `db:"actual"`
	Expected string `db:"expected"`
}

// Save implements Collector.
func (li *Lite) Save(ctx context.Context, titleOrExpected, actual string) (int64, error) {
	nowUTC := time.Now().UTC()
	id := nowUTC.UnixNano()
	_, err := li.db.ExecContext(ctx, `
		INSERT INTO issues (id, expected, actual, created_at)
		VALUES (?, ?, ?, ?)
	`, id, titleOrExpected, actual, nowUTC.Format(time.RFC3339))
	return id, err
}

func getVersion(db *sqlx.DB) (string, error) {
	var s string
	err := db.Get(&s, `SELECT key FROM issue_version`)
	return s, err
}

func migrateSchema(db *sqlx.DB) error {
	schemas := []string{
		dbSchemaTableVersion,
		dbSchemaInsertVersion,
		dbSchemaTableIssues,
	}
	for i, s := range schemas {
		_, err := db.Exec(s)
		if err != nil {
			return fmt.Errorf("migrate schema at: %d %s", i, err)
		}
	}
	return nil
}

const (
	dbSchemaTableVersion = `
		CREATE TABLE issue_version (
			key TEXT PRIMARY KEY
		);
	`
	dbSchemaInsertVersion = `
		INSERT INTO issue_version (key) VALUES ('0.0.1');
	`
	dbSchemaTableIssues = `
		CREATE TABLE issues (
			id INTEGER PRIMARY KEY,
			expected TEXT NOT NULL,
			actual TEXT NOT NULL,
			created_at TEXT NOT NULL
		);
	`
)
