package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"

	"github.com/0xf6i/cli/internal/config"
	sshserver "github.com/0xf6i/cli/internal/ssh"
)

func main() {
	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		cfgPath = "config.yml"
	}

	cfg, err := config.Load(cfgPath)
	if err != nil {
		log.Warn("Could not load config, using defaults", "err", err)
		cfg, _ = config.Load("/dev/null") // triggers defaults
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := sshserver.Start(cfg); err != nil {
			log.Fatal("Server error", "err", err)
		}
	}()

	<-done
	log.Info("Server stopped")
}
