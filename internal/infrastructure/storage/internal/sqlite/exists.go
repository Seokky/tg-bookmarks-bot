package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
)

func (s *Storage) Exists(username, bookmark string) (bool, error) {
	query := "SELECT bookmark FROM " + tableName + " WHERE username = ? AND bookmark = ?"

	var existing string

	if err := s.db.QueryRow(query, username, bookmark).Scan(&existing); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("can't check if exists: %w", err)
	}

	return len(existing) > 0, nil
}
