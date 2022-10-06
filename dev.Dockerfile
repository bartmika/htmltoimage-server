# The base go-image
FROM golang:1.18-alpine

# Special Thanks: https://stackoverflow.com/a/59139842
# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
RUN apk update && apk upgrade && apk add --no-cache bash git && apk add --no-cache chromium

# Installs latest Chromium package.
RUN echo @edge http://nl.alpinelinux.org/alpine/edge/community >> /etc/apk/repositories \
    && echo @edge http://nl.alpinelinux.org/alpine/edge/main >> /etc/apk/repositories \
    && apk add --no-cache \
    harfbuzz@edge \
    nss@edge \
    freetype@edge \
    ttf-freefont@edge \
    && rm -rf /var/cache/* \
    && mkdir /var/cache/apk

CMD chromium-browser --headless --disable-gpu --remote-debugging-port=9222 --disable-web-security --safebrowsing-disable-auto-update --disable-sync --disable-default-apps --hide-scrollbars --metrics-recording-only --mute-audio --no-first-run --no-sandbox
# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

COPY . /go/src/github.com/bartmika/htmltoimage-server
WORKDIR /go/src/github.com/bartmika/htmltoimage-server

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy only `.go` files, if you want all files to be copied then replace `with `COPY . .` for the code below.
COPY *.go .

# Install our third-party application for hot-reloading capability.
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "install", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -polling -log-prefix=false -build="go build ." -command="./htmltoimage-server serve" -directory="./"

# BUILD
# docker build --rm -t htmltoimage-server -f dev.Dockerfile .

# EXECUTE
# docker run -d -p 8000:8000 htmltoimage-server
