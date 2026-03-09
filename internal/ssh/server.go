package ssh

import (
	"github.com/0xf6i/cli/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
)

func Start(host, port string) error {
	srv, err := wish.NewServer(
		wish.WithAddress(host+":"+port),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
		),
	)
	if err != nil {
		return err
	}

	log.Info("Starting SSH server", "host", host, "port", port)
	return srv.ListenAndServe()
}

func teaHandler(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
	return ui.NewModel(), []tea.ProgramOption{tea.WithAltScreen()}
}
