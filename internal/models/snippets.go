package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expired time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	return 0, nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
