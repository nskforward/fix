package fix

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"net"
)

type Conn struct {
	rwc io.ReadWriteCloser
}

func NewConn(network, addr string) (*Conn, error) {
	rwc, err := net.Dial(network, addr)
	return &Conn{rwc: rwc}, err
}

func NewConnTLS(network, addr string, cert []byte) (*Conn, error) {
	roots := x509.NewCertPool()
	if cert != nil {
		ok := roots.AppendCertsFromPEM(cert)
		if !ok {
			return nil, errors.New("cannot parse PEM certificate for tcp dial")
		}
	}
	rwc, err := tls.Dial(network, addr, &tls.Config{
		RootCAs:            roots,
		InsecureSkipVerify: true,
	})
	return &Conn{rwc: rwc}, err
}

func (conn *Conn) Write(p []byte) (int, error) {
	return conn.rwc.Write(p)
}

func (conn *Conn) Close() {
	_ = conn.rwc.Close()
}

func (conn *Conn) Read(p []byte) (int, error) {
	return conn.rwc.Read(p)
}
