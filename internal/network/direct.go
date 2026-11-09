package network

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

func NewStream(ctx context.Context, h host.Host, multiaddr string) (network.Stream, error) {
	peerAddr, err := peer.AddrInfoFromString(multiaddr)
	if err != nil {
		return nil, fmt.Errorf("Error building AddrInfo: %v", err)
	}

	stream, err := h.NewStream(ctx, peerAddr.ID, "/arkdy/1.0.0")

	if err != nil {
		return nil, fmt.Errorf("Error opening a stream: %v", err)
	}

	// rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	return stream, nil
}
