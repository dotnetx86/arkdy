package network

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/dotnetx86/arkdy/internal/identity"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	ma "github.com/multiformats/go-multiaddr"
)

func NewHost(port int) (host.Host, error) {
	addr, _ := ma.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))

	priv, err := identity.CreateIdentity()
	if err != nil {
		return nil, err
	}

	h, err := libp2p.New(
		libp2p.Identity(priv),
		libp2p.ListenAddrs(addr),
		libp2p.Security(noise.ID, noise.New),
	)
	if err != nil {
		return nil, fmt.Errorf("Error creating a host: %v", err)
	}

	slog.Debug("Successfully created a host", "priv", priv, "addr", addr)
	return h, nil
}

func Connect(h host.Host, multiaddr string) error {
	peerAddr, err := peer.AddrInfoFromString(multiaddr)
	if err != nil {
		return fmt.Errorf("Error building AddrInfo: %v", err)
	}

	err = h.Connect(context.Background(), *peerAddr)
	if err != nil {
		return fmt.Errorf("Error connecting to a peer: %v", err)
	}

	slog.Debug("Successfully connected to the peer", "multiaddr", multiaddr)
	return nil
}
