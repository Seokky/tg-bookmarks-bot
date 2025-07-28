package sqlite

import (
	"fmt"
)

// Insert performs inserting user's bookmark into database table
func (s *Storage) Insert(username, bookmark string) error {
	query := "INSERT INTO " + tableName + " (username, bookmark) VALUES (?, ?)"

	_, err := s.db.Exec(query, username, bookmark)
	if err != nil {
		return fmt.Errorf("failed to insert record: %w", err)
	}

	return nil
}
