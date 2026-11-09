package cli

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/dotnetx86/arkdy/client"
	"github.com/dotnetx86/arkdy/internal/messaging"
)

func Run(ctx context.Context) {
	dest := flag.String("d", "", "Destination address")
	port := flag.Int("p", 9000, "Port number")
	help := flag.Bool("help", false, "Display the help message")
	debug := flag.Bool("debug", false, "Show debug information")

	flag.Parse()

	if *help {
		fmt.Printf("This program demonstrates a simple p2p chat application using libp2p\n\n")
		fmt.Println("Usage: Run './chat -d <MULTIADDR>' where <MULTIADDR> is multiaddress of previous listener host.")

		os.Exit(0)
	}

	level := slog.LevelInfo
	if *debug {
		level = slog.LevelDebug
	}

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})
	slog.SetDefault(slog.New(handler))

	client.Run(ctx, *dest, *port, PutData, GetData)

}

func PutData(rw *bufio.ReadWriter) error {
	for {
		str, err := messaging.Read(rw)
		// str, _ := rw.ReadString('\n')

		if err != nil {
			return err
		}
		if str != "\n" {
			// Green console colour: 	\x1b[32m
			// Reset console colour: 	\x1b[0m
			fmt.Printf("%s", str)
		}

	}
}

func GetData(rw *bufio.ReadWriter) error {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			return err
		}

		messaging.Write(rw, sendData)

		// rw.WriteString(fmt.Sprintf("%s\n", sendData))
		// rw.Flush()
	}
}
