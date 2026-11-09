package messaging

import (
	"bufio"
	"context"
	"fmt"
	"log"

	"github.com/dotnetx86/arkdy/internal/network"
	"github.com/libp2p/go-libp2p/core/host"
)

func Init(ctx context.Context, port int) (host.Host, error) {
	h, err := network.NewHost(port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Run './chat -d /ip4/127.0.0.1/tcp/%v/p2p/%s' on another console.\n", port, h.ID())

	return h, err
}

func Read(rw *bufio.ReadWriter) (string, error) {
	str, err := rw.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("Error reading from a stream: %v", err)
	}

	return str, nil
}

func Write(rw *bufio.ReadWriter, sendData string) error {
	_, err := rw.WriteString(fmt.Sprintf("%s\n", sendData))
	if err != nil {
		return fmt.Errorf("Error writing to a stream: %v", err)
	}

	err = rw.Flush()
	if err != nil {
		return fmt.Errorf("Error flushing the writer: %v", err)
	}

	return nil
}
