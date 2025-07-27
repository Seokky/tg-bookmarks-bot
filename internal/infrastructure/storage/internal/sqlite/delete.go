package sqlite

import "fmt"

func (s *Storage) Delete(username, bookmark string) error {
	query := "DELETE FROM " + tableName + " WHERE username = ? AND bookmark = ?"

	_, err := s.db.Exec(query, username, bookmark)
	if err != nil {
		return fmt.Errorf("failed to delete record: %w", err)
	}

	return nil
}
