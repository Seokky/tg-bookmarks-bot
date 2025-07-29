# tg-bookmarks-bot

## About
This is implementation of Telegram bot, that can store all of my bookmarks on useful articles and videos and provides ability to pick the random one of stored bookmarks to explore it. 

Available at @bookmarks_storer_bot in Telegram. To use this, contact me (@sykkoes) so that I can add you to the whitelist. But better way for you is independently deploy it on your VPS.

## Commands
- /random, /r - get random one of user stored bookmarks (bookmark will be deleted from storage).
- /count, /c - get count of user's bookmarks
- just text or link - performs save provided message as bookmark

## Implementation
Bot using long polling to get updates, so everything needed to make it work - is running process of this program. SQLite used as storage.

Scheme is a little bit outdated but in general is relevant still:

<img width="696" height="711" alt="image" src="https://github.com/user-attachments/assets/2675f18c-4653-4795-b348-27c05a70348a" />
