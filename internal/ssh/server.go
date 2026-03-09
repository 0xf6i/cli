package ssh

import (
	"github.com/0xf6i/cli/internal/config"
	"github.com/0xf6i/cli/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
)

func Start(cfg config.Config) error {
	host := cfg.Server.Host
	port := cfg.Server.Port

	srv, err := wish.NewServer(
		wish.WithAddress(host+":"+port),
		wish.WithMiddleware(
			bubbletea.Middleware(func(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
				return ui.NewModel(cfg), []tea.ProgramOption{tea.WithAltScreen()}
			}),
		),
	)
	if err != nil {
		return err
	}

	log.Info("Starting SSH server", "host", host, "port", port)
	return srv.ListenAndServe()
}
