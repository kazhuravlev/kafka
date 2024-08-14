package admin

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"net"
	"strconv"
	"time"
)

func (c *Client) getConnection(ctx context.Context) (*kafka.Conn, error) {
	dialer := &kafka.Dialer{ //nolint:exhaustruct
		Timeout:   10 * time.Second, //nolint:gomnd
		DualStack: true,
		TLS:       c.opts.tlsConfig,
	}

	conn, err := dialer.DialContext(ctx, "tcp", c.opts.bootstrapServers[0])
	if err != nil {
		return nil, fmt.Errorf("dial kafka: %w", err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return nil, fmt.Errorf("get cluster controller: %w", err)
	}

	controllerAddr := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))

	controllerConn, err := dialer.DialContext(ctx, "tcp", controllerAddr)
	if err != nil {
		return nil, fmt.Errorf("dial controller: %w", err)
	}

	go func() {
		<-ctx.Done()
		controllerConn.Close()
	}()

	return controllerConn, nil
}
