package ch00

import (
	"advanced/ch00/server"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServerBuilder(t *testing.T) {
	srv, err := new(server.ServerBuilder).
		New("127.0.0.1", -1).
		WithProtocol("xxx").
		WithMaxConn(1024).
		WithTimeout(30 * time.Second).
		Build()
	assert.Error(t, err)
	assert.Equal(t, "127.0.0.1", srv.Addr)
	assert.Equal(t, 1024, srv.MaxConn)
	assert.Equal(t, 30*time.Second, srv.Timeout)
	assert.Equal(t, true, errors.Is(err, server.ErrInvalidPort))
	assert.Equal(t, false, errors.Is(err, server.ErrInvalidAddress))
	assert.Equal(t, true, errors.Is(err, server.ErrInvalidProtocol))
	assert.Equal(t, false, errors.Is(err, server.ErrInvalidMaxConn))
	assert.Equal(t, false, errors.Is(err, server.ErrInvalidTimeout))
}
