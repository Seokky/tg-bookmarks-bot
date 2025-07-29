package storer

import "fmt"

// Random performs getting the random one of the user's bookmarks
func (s *Storer) Random(username string) (bookmark string, err error) {
	bookmark, err = s.storage.Random(username)
	if err != nil {
		return "", err
	}

	defer func() {
		if bookmark == "" || err != nil {
			return
		}

		err = s.storage.Delete(username, bookmark)
		if err != nil {
			bookmark = ""
			err = fmt.Errorf("can't delete bookmark: %w", err)
		}
	}()

	return bookmark, err
}
