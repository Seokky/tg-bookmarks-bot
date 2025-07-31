# tg-bookmarks-bot

## About
This is implementation of Telegram bot, that can store all of my bookmarks on useful articles and videos and provides ability to pick the random one of stored bookmarks to explore it. 

Available at @bookmarks_storer_bot in Telegram. To use this, contact me (@sykkoes) so that I can add you to the whitelist. But better way for you is independently deploy it on your VPS.

## Commands
- /random, /r - get random one of user stored bookmarks (bookmark will be deleted from storage).
- /count, /c - get count of user's bookmarks
- just text or link - performs save provided message as bookmark

## Implementation
Bot using long polling to get updates, so everything needed to make it work is running process of this program.

## To do
- use context to add timeout to poll request
- use context to add timeout to send request
- implement tests for net.Fetch using httptest
- implement test for net.Send using httptest
- pass to net.Fetch url instead of entrypoint
- pass to net.Send url instead of entrypoint
- synchronize order of updates handling for concrete user