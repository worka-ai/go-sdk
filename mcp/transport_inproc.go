// Copyright 2025 The Go MCP SDK Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mcp

import (
	"context"
	"sync"

	"github.com/modelcontextprotocol/go-sdk/jsonrpc"
)

// An InProcTransport is a [Transport] that communicates over in-process
// channels without JSON framing.
//
// InProcTransports should be constructed using [NewInProcTransports], which
// returns two transports connected to each other.
type InProcTransport struct {
	pair     *inProcPair
	incoming <-chan jsonrpc.Message
	outgoing chan<- jsonrpc.Message
	isServer bool
}

// Connect implements the [Transport] interface.
func (t *InProcTransport) Connect(context.Context) (Connection, error) {
	if t.isServer {
		return &inProcServerConn{
			pair:     t.pair,
			incoming: t.incoming,
			outgoing: t.outgoing,
		}, nil
	}
	return &inProcClientConn{
		pair:     t.pair,
		incoming: t.incoming,
		outgoing: t.outgoing,
	}, nil
}

// NewInProcTransports returns two [InProcTransport] objects that connect
// to each other.
//
// The resulting transports are symmetrical: use either to connect to a server,
// and then the other to connect to a client. Servers must be connected before
// clients, as the client initializes the MCP session during connection.
func NewInProcTransports() (*InProcTransport, *InProcTransport) {
	pair := &inProcPair{
		aToB:   make(chan jsonrpc.Message, 16),
		bToA:   make(chan jsonrpc.Message, 16),
		closed: make(chan struct{}),
	}
	return &InProcTransport{
			pair:     pair,
			incoming: pair.aToB,
			outgoing: pair.bToA,
			isServer: true,
		}, &InProcTransport{
			pair:     pair,
			incoming: pair.bToA,
			outgoing: pair.aToB,
			isServer: false,
		}
}

type inProcPair struct {
	aToB chan jsonrpc.Message
	bToA chan jsonrpc.Message

	closed    chan struct{}
	closeOnce sync.Once
}

func (p *inProcPair) close() {
	p.closeOnce.Do(func() {
		close(p.closed)
		close(p.aToB)
		close(p.bToA)
	})
}

type inProcServerConn struct {
	pair     *inProcPair
	incoming <-chan jsonrpc.Message
	outgoing chan<- jsonrpc.Message
}

func (c *inProcServerConn) SessionID() string { return "" }

func (c *inProcServerConn) Read(ctx context.Context) (jsonrpc.Message, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-c.pair.closed:
		return nil, ErrConnectionClosed
	case msg, ok := <-c.incoming:
		if !ok {
			return nil, ErrConnectionClosed
		}
		return msg, nil
	}
}

func (c *inProcServerConn) Write(ctx context.Context, msg jsonrpc.Message) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-c.pair.closed:
		return ErrConnectionClosed
	case c.outgoing <- msg:
		return nil
	}
}

func (c *inProcServerConn) Close() error {
	c.pair.close()
	return nil
}

func (c *inProcServerConn) sessionUpdated(ServerSessionState) {}

type inProcClientConn struct {
	pair     *inProcPair
	incoming <-chan jsonrpc.Message
	outgoing chan<- jsonrpc.Message
}

func (c *inProcClientConn) SessionID() string { return "" }

func (c *inProcClientConn) Read(ctx context.Context) (jsonrpc.Message, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-c.pair.closed:
		return nil, ErrConnectionClosed
	case msg, ok := <-c.incoming:
		if !ok {
			return nil, ErrConnectionClosed
		}
		return msg, nil
	}
}

func (c *inProcClientConn) Write(ctx context.Context, msg jsonrpc.Message) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-c.pair.closed:
		return ErrConnectionClosed
	case c.outgoing <- msg:
		return nil
	}
}

func (c *inProcClientConn) Close() error {
	c.pair.close()
	return nil
}

func (c *inProcClientConn) sessionUpdated(state clientSessionState) {}
