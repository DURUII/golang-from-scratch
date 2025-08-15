package server

import (
	"errors"
	"net"
	"time"
)

var (
	ErrInvalidAddress  = errors.New("address cannot be empty or invalid")
	ErrInvalidPort     = errors.New("invalid port number, must be between 1 and 65535")
	ErrInvalidProtocol = errors.New("invalid protocol, only 'tcp' or 'udp' are allowed")
	ErrInvalidMaxConn  = errors.New("maxConn must be greater than zero")
	ErrInvalidTimeout  = errors.New("timeout must be greater than zero")
)

type ServerBuilder struct {
	Server
}

func (sb *ServerBuilder) New(addr string, port int) *ServerBuilder {
	if net.ParseIP(addr) == nil {
		_, err := net.ResolveIPAddr("ip", addr)
		if err != nil {
			sb.Err = errors.Join(sb.Err, ErrInvalidAddress)
		}
	}

	if port < 0 || port > 65535 {
		sb.Err = errors.Join(sb.Err, ErrInvalidPort)
	}

	sb.Addr = addr
	sb.Port = port
	return sb
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	if protocol != "tcp" && protocol != "udp" {
		sb.Err = errors.Join(sb.Err, ErrInvalidProtocol)
		return sb
	}
	sb.Protocol = protocol
	return sb
}

func (sb *ServerBuilder) WithMaxConn(maxConn int) *ServerBuilder {
	if maxConn < 0 {
		sb.Err = errors.Join(sb.Err, ErrInvalidMaxConn)
	}
	sb.MaxConn = maxConn
	return sb
}

func (sb *ServerBuilder) WithTimeout(timeout time.Duration) *ServerBuilder {
	if timeout <= 0 {
		sb.Err = errors.Join(sb.Err, ErrInvalidTimeout)
	}
	sb.Timeout = timeout
	return sb
}

func (sb *ServerBuilder) Build() (*Server, error) {
	return &sb.Server, sb.Err
}
