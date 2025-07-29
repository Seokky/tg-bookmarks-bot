package storer

// Count performs getting count of user's bookmarks
func (s *Storer) Count(username string) (count int, err error) {
	return s.storage.Count(username)
}
