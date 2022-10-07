# HTML to Image Server

Self-hosted HTML to image service powered by Golang RPC server for your web-applications.

## Installation

Create the `docker-compose.yml` file your directory.

```yaml
version: '3.8'
services:
  app:
    container_name: htmltoimage_server_app
    image: bartmika/htmltoimage-server:latest
    stdin_open: true
    environment:
        HTMLTOIMAGE_SERVER_IP: 0.0.0.0
        HTMLTOIMAGE_SERVER_PORT: 8002
        HTMLTOIMAGE_SERVER_CHROME_HEADLESS_WS_URL: ws://browserless:9090
    restart: unless-stopped
    ports:
      - "8002:8002"
    depends_on:
      - browserless
    links:
      - browserless

  browserless:
    image: browserless/chrome:latest
    environment:
      - DEBUG=browserless:*
      - MAX_CONCURRENT_SESSIONS=10
      - CONNECTION_TIMEOUT=60000
      - MAX_QUEUE_LENGTH=20
      - PREBOOT_CHROME=true
      - DEMO_MODE=false
      - HOST=0.0.0.0
      - ENABLE_DEBUGGER=false
      - PORT=9090
      - WORKSPACE_DELETE_EXPIRED=true
    container_name: "htmltoimage_server_browserless"
    restart: unless-stopped
```

Start the service.

```shell
$ docker compose up -d
```

## Usage

Here is a sample program of how to utilize the service.

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/bartmika/htmltoimage-server/pkg/rpc"
)

// NOTE:
// Before you begin running this code, make sure the HTML to Image server is
// running before you execute this program.

func main() {
	// Generate the HTML to Image server based on your `docker-compose.yml`.
	serverIP := "127.0.0.1"
	serverPort := "8002"
	applicationAddress := fmt.Sprintf("%s:%s", serverIP, serverPort)

	// Connect to a running client.
	client, err := rpc.NewClient(applicationAddress, 3, 15*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the remote call to the HTML to Image server.
	res, err := client.Screenshot("https://brank.as/")
	if err != nil {
		log.Fatal("Sample command failed generating image with error:", err)
	}

	// Save our file.
	if err := ioutil.WriteFile("data/"+res.FileName+".png", res.Content, 0o644); err != nil {
		log.Fatal("Sample command failed writing file with error:", err)
	}
	log.Println("Saved file:", res.FileName)
}
```

Finally if you look in your `data` folder, you'll see the image downloaded!
