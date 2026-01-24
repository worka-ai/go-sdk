package mcp

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/internal/jsonrpc2"
	"github.com/modelcontextprotocol/go-sdk/jsonrpc"
)

func TestInProcTransportRoundTrip(t *testing.T) {
	serverTransport, clientTransport := NewInProcTransports()

	serverConn, err := serverTransport.Connect(context.Background())
	if err != nil {
		t.Fatalf("server connect: %v", err)
	}
	clientConn, err := clientTransport.Connect(context.Background())
	if err != nil {
		t.Fatalf("client connect: %v", err)
	}
	defer serverConn.Close()
	defer clientConn.Close()

	go func() {
		_ = clientConn.Write(context.Background(), &jsonrpc.Request{
			ID:     jsonrpc2.Int64ID(1),
			Method: "ping",
		})
	}()

	msg, err := serverConn.Read(context.Background())
	if err != nil {
		t.Fatalf("server read: %v", err)
	}
	req, ok := msg.(*jsonrpc.Request)
	if !ok {
		t.Fatalf("expected request, got %T", msg)
	}
	if req.Method != "ping" {
		t.Fatalf("expected method ping, got %q", req.Method)
	}
}
