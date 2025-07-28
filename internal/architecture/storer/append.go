package storer

import "errors"

var (
	ErrBookmarkAlreadyExists = errors.New("bookmark already exists")
)

func (s *Storer) Append(username, bookmark string) error {
	exists, err := s.storage.Exists(username, bookmark)
	if err != nil {
		return err
	}

	if exists {
		return ErrBookmarkAlreadyExists
	}

	return s.storage.Insert(username, bookmark)
}
