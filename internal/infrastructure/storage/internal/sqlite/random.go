package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
)

func (s *Storage) Random(username string) (string, error) {
	query := "SELECT bookmark FROM " + tableName + ` WHERE username = ? ORDER BY random() LIMIT 1`

	var bookmark string

	if err := s.db.QueryRow(query, username).Scan(&bookmark); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}

		return "", fmt.Errorf("can't select random: %w", err)
	}

	return bookmark, nil
}
