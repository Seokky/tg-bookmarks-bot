// Package commander
package commander

import (
	"strings"
	storerPkg "tg-bookmarks-bot/internal/architecture/storer"
)

func ProcessCommand(storer *storerPkg.Storer, text, username string) string {
	text = strings.TrimSpace(text)

	switch text {
	case "":
		return "Incorrect command"
	case "/random":
		bookmark, err := storer.Random(username)
		if err != nil {
			return "Cannot get random bookmark. Please, contact @sykkoes"
		}
		if bookmark == "" {
			return "You have no bookmark yet"
		}
		// TODO go storer.Delete(username, bookmark)
		return bookmark
	default:
		// TODO handle if bookmark already exists
		err := storer.Append(username, text)
		if err != nil {
			return "Cannot append bookmark. Please, contact @sykkoes"
		}
		return "Bookmark appended successfully"
	}
}
