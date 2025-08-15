package ch00

import (
	"advanced/ch00/config"
	"advanced/ch00/server"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
	"time"
)

type ServerOption func(*server.Server)

func NewServer(addr string, port int, options ...ServerOption) *server.Server {
	srv := &server.Server{
		Addr: addr,
		Port: port,
	}
	for _, option := range options {
		option(srv)
	}
	return srv
}

// WithProtocol is an option to set the protocol of the server.
func WithProtocol(protocol string) ServerOption {
	return func(s *server.Server) {
		if protocol != "tcp" && protocol != "udp" {
			panic(server.ErrInvalidProtocol)
		}
		s.Protocol = protocol
	}
}

// WithMaxConn is an option to set the max connections of the server.
func WithMaxConn(maxConn int) ServerOption {
	return func(s *server.Server) {
		if maxConn <= 0 {
			panic(server.ErrInvalidMaxConn)
		}
		s.MaxConn = maxConn
	}
}

// WithTimeout is an option to set the timeout of the server.
func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *server.Server) {
		if timeout <= 0 {
			panic(server.ErrInvalidTimeout)
		}
		s.Timeout = timeout
	}
}

func BuildSrv(opts ...ServerOption) (srv *server.Server, err error) {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("recovered from: ", "panic", r)
			err = r.(error)
			return
		}
	}()

	srv = NewServer(config.GetString("server.addr"), config.GetInt("server.port"))
	for _, opt := range opts {
		opt(srv)
	}
	return
}

func TestServerOptions(t *testing.T) {
	srv, _ := BuildSrv(
		WithProtocol(config.GetString("server.extra.protocol")),
		WithMaxConn(config.GetInt("server.extra.max_conn")),
		WithTimeout(time.Duration(config.GetInt("server.extra.timeout"))*time.Millisecond),
	)
	// in this case, you will see:
	// invalid memory address or nil pointer dereference [recovered]
	assert.Equal(t, "tcp", srv.Protocol)
	assert.Equal(t, 1000, srv.MaxConn)
	assert.Equal(t, 100*time.Millisecond, srv.Timeout)
}
