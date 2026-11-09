package client

import (
	"bufio"
	"context"
	"log"

	"github.com/dotnetx86/arkdy/internal/messaging"
	"github.com/dotnetx86/arkdy/internal/network"
)

func Run(ctx context.Context, dest string, port int, PutData func(*bufio.ReadWriter) error, GetData func(*bufio.ReadWriter) error) {
	h, err := messaging.Init(ctx, port)
	if err != nil {
		log.Fatal(err)
	}

	if dest != "" {
		err := network.Connect(h, dest)
		if err != nil {
			log.Fatal(err)
		}

		stream, err := network.NewStream(ctx, h, dest)
		if err != nil {
			log.Fatal(err)
		}

		rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

		go PutData(rw)
		go GetData(rw)

	}

}
