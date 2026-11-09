package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/dotnetx86/arkdy/internal/cli"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cli.Run(ctx)

	<-ctx.Done()

}
