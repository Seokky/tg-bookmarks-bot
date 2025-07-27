// Package commander
package commander

import (
	"errors"
	"fmt"
	"log"
	"strings"
	storerPkg "tg-bookmarks-bot/internal/architecture/storer"
)

const (
	commandEmpty            = ""
	commandGetRandom        = "/random"
	commandGetRandomShortly = "/r"
	commandPop              = "/pop"
	commandCount            = "/count"
	commandCountShortly     = "/c"
)

func ProcessCommand(storer *storerPkg.Storer, text, username string) (message string) {
	defer func() {
		if err := recover(); err != nil {
			logger := log.Default()
			logger.Println(err)
			message = "Unhandled error. Please, contact @sykkoes"
		}
	}()

	switch text = strings.TrimSpace(text); text {
	case commandEmpty:
		return "Incorrect command"
	case commandGetRandom, commandGetRandomShortly:
		bookmark, err := storer.Random(username)
		if err != nil {
			return "Cannot get random bookmark. Please, contact @sykkoes"
		}
		if bookmark == "" {
			return "You have no bookmark yet"
		}
		return bookmark
	case commandCount, commandCountShortly:
		count, err := storer.Count(username)
		if err != nil {
			return "Cannot get bookmark count. Please, contact @sykkoes"
		}
		return fmt.Sprintf("You have %d bookmarks", count)
	default:
		err := storer.Append(username, text)
		if err != nil {
			if errors.Is(err, storerPkg.ErrBookmarkAlreadyExists) {
				return "Bookmark already exists"
			}
			return "Cannot append bookmark. Please, contact @sykkoes"
		}
		return "Bookmark appended successfully"
	}
}
