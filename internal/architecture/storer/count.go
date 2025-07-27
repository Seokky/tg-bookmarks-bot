package storer

func (s *Storer) Count(username string) (count int, err error) {
	return s.storage.Count(username)
}
