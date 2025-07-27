package storer

func (s *Storer) Random(username string) (string, error) {
	return s.storage.Random(username)
}
