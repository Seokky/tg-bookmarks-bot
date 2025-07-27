package storer

func (s *Storer) Append(username, bookmark string) error {
	return s.storage.Insert(username, bookmark)
}
