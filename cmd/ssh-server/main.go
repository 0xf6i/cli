package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"

	sshserver "github.com/0xf6i/cli/internal/ssh"
)

func main() {
	host := os.Getenv("SSH_HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("SSH_PORT")
	if port == "" {
		port = "2222"
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := sshserver.Start(host, port); err != nil {
			log.Fatal("Server error", "err", err)
		}
	}()

	<-done
	log.Info("Server stopped")
}
