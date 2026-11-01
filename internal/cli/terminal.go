package cli

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
)

func Run(ctx context.Context) {
	dest := flag.String("d", "", "Destination address")
	help := flag.Bool("help", false, "Display the help message")

	flag.Parse()

}

func ReadData(rw *bufio.ReadWriter) {
	for {
		str, _ := rw.ReadString('\n')

		if str == "" {
			return
		}
		if str != "\n" {
			// Green console colour: 	\x1b[32m
			// Reset console colour: 	\x1b[0m
			fmt.Printf("%s", str)
		}

	}
}

func WriteData(rw *bufio.ReadWriter) {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		rw.WriteString(fmt.Sprintf("%s\n", sendData))
		rw.Flush()
	}
}
