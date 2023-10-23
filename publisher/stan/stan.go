package stan

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

type ClientConn struct {
	SC stan.Conn
}

func NewStanClient(clusterId, clientId, url string) (*ClientConn, error) {
	sc, err := stan.Connect(clusterId, clientId, stan.NatsURL(url))
	if err != nil {
		return nil, fmt.Errorf("conn: %w", err)
	}

	cc := &ClientConn{
		SC: sc,
	}

	return cc, nil
}

func (cc *ClientConn) Close() error {
	err := cc.SC.Close()
	if err != nil {
		return fmt.Errorf("close: %w", err)
	}
	return nil
}
