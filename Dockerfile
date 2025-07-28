FROM golang:1.24

WORKDIR /usr/src/tgbotbookmarks

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

RUN mkdir /usr/local/bin/tgbotbookmarks

COPY . .
RUN go build -v -o /usr/local/bin/tgbotbookmarks ./...

CMD ["/usr/local/bin/tgbotbookmarks/tg-bookmarks-bot"]