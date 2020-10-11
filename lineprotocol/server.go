package lineprotocol

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// ErrServerClosed occurs when a tcp server is closed.
var ErrServerClosed = errors.New("Server closed")

// Proto represents the type of connection
type Proto int

const (
	NONE Proto = iota
	MQTT
	GRPC
)

//Handler is a callback which get called when a tcp, websocket connection is established or a grpc stream is established
type Handler func(c net.Conn, proto Proto)

type options struct {
	TLSConfig *tls.Config
	KeepAlive bool
}

// Options it contains configurable options for client
type Options interface {
	set(*options)
}

// fOption wraps a function that modifies options into an
// implementation of the Option interface.
type fOption struct {
	f func(*options)
}

func (fo *fOption) set(o *options) {
	fo.f(o)
}

func newFuncOption(f func(*options)) *fOption {
	return &fOption{
		f: f,
	}
}

// WithDefaultOptions will create client connection with some default values.
//   KeepAlive: true
//   TlsConfig: nil
func WithDefaultOptions() Options {
	return newFuncOption(func(o *options) {
		o.KeepAlive = true
		o.TLSConfig = nil
	})
}

// WithTLSConfig will set an SSL/TLS configuration to be used when connecting
// to server.
func WithTLSConfig(t *tls.Config) Options {
	return newFuncOption(func(o *options) {
		o.TLSConfig = t
	})
}

type Server interface {
	// Serve serve the requests if type tcp, websocket or grpc stream
	Serve(net.Listener) error
}

type server struct {
	sync.Mutex
	opts    *options
	Handler Handler //The handler to invoke when a connection is accepted
}

func signalHandler() <-chan bool {
	stop := make(chan bool)

	signchan := make(chan os.Signal, 1)
	signal.Notify(signchan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		// Wait for a signal. Don't care which signal it is
		sig := <-signchan
		fmt.Printf("Signal received: '%s', shutting down\n", sig)
		stop <- true
	}()

	return stop
}
