package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
)

func (s *Storage) Count(username string) (int, error) {
	query := "SELECT COUNT(*) FROM " + tableName + ` WHERE username = ?`

	var count int

	if err := s.db.QueryRow(query, username).Scan(&count); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}

		return 0, fmt.Errorf("can't select count: %w", err)
	}

	return count, nil
}
