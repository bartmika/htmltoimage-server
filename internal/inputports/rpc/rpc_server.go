package rpc

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bartmika/htmltoimage-server/internal/app"
	"github.com/bartmika/htmltoimage-server/internal/config"
)

type RPC struct {
	App         app.HTMLToImageApp
	tcpListener *net.TCPListener
}

func NewServer(appConfig *config.Conf, a app.HTMLToImageApp) *RPC {

	applicationAddress := fmt.Sprintf("%s:%s", appConfig.Server.IP, appConfig.Server.Port)
	log.Println("Initializing address for", applicationAddress)

	tcpAddr, err := net.ResolveTCPAddr("tcp", applicationAddress)
	if err != nil {
		log.Fatal(err)
	}

	rpcServer := &RPC{App: a}

	// Attach the
	rpc.Register(rpcServer)
	rpc.HandleHTTP()

	log.Println("RPC API was initialized.")
	l, e := net.ListenTCP("tcp", tcpAddr)
	if e != nil {
		l.Close()
		log.Fatal("RPC API failed to initialize:", e.Error())
	}

	log.Println("RPC API is listening now.")
	rpcServer.tcpListener = l
	return rpcServer
}

func (rpcServer *RPC) RunMainRuntimeLoop() error {
	log.Println("RPC API is starting now...")

	// The following code will attach a background handler to run when the
	// application detects a shutdown signal.
	// Special thanks via https://guzalexander.com/2017/05/31/gracefully-exit-server-in-go.html
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs // Block execution until signal from terminal gets triggered here.

		// Finish any RPC communication taking place at the moment before
		// shutting down the RPC server.
		rpcServer.tcpListener.Close()
	}()

	// Attach the following anonymous function to run on all cases (ex: panic,
	// termination signal, etc) so we can gracefully shutdown the service.
	defer func() {
		rpcServer.stopRuntimeLoop()
	}()

	// Safety net for 'too many open files' issue on legacy code.
	// Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	// Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	http.DefaultClient.Timeout = time.Minute * 10

	// DEVELOPER NOTES:
	// If you get "too many open files" then please read the following article
	// http://publib.boulder.ibm.com/httpserv/ihsdiag/too_many_open_files.html
	// so you can run in your console:
	// $ ulimit -H -n 4096
	// $ ulimit -n 4096

	log.Println("RPC API is ready and running.")

	// Run the main loop blocking code.
	http.Serve(rpcServer.tcpListener, nil)

	return nil
}

func (rpcServer *RPC) stopRuntimeLoop() error {
	log.Printf("Starting graceful shutdown now...")
	rpcServer.tcpListener.Close()
	log.Printf("Terminated TCPListener.")
	log.Printf("Graceful shutdown finished.")
	return nil
}
