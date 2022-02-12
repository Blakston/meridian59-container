package maintenance

import (
	"bytes"
	"net"
	"time"
)

// Handler ...
type Handler struct {
	con *net.TCPConn
	err error
}

// Close ...
func (a *Handler) Close() {
	if a.err != nil {
		return
	}
	a.con.Close()
}

// Connect ...
func (a *Handler) Connect(addr string) {
	if a.err != nil {
		return
	}
	// connect to the maintenance port
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		a.err = err
		return
	}
	con, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		a.err = err
		return
	}
	a.con = con
}

// Error ...
func (a *Handler) Error() error {
	return a.err
}

// Receive ...
func (a *Handler) Receive() (out string) {
	if a.err != nil {
		return
	}
	var result bytes.Buffer
	go func() {
		for {
			reply := make([]byte, 8192)
			if _, err := a.con.Read(reply); err != nil {
				a.err = err
				break
			}
			result.WriteString(string(reply))
		}
	}()
	<-time.After(3 * time.Second)
	return result.String()
}

// Send ...
func (a *Handler) Send(in string) {
	if a.err != nil {
		return
	}
	if _, err := a.con.Write([]byte(in + "\r\n")); err != nil {
		a.err = err
		return
	}
}

// NewHandler creates a new maintenance handler.
func NewHandler() *Handler {
	return &Handler{}
}
