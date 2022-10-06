# The base go-image
FROM golang:1.18

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
