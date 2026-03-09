package ui

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Tick messages — distinct types to avoid collision.

type animationTickMsg struct{}

func animationTick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(time.Time) tea.Msg {
		return animationTickMsg{}
	})
}

type typingTickMsg struct{}

func typingTick() tea.Cmd {
	return tea.Tick(35*time.Millisecond, func(time.Time) tea.Msg {
		return typingTickMsg{}
	})
}

// handleCommand parses user input and transitions state.
func (m Model) handleCommand(input string) (Model, tea.Cmd) {
	raw := strings.ToLower(strings.TrimPrefix(input, "/"))
	parts := strings.Fields(raw)
	if len(parts) == 0 {
		return m, nil
	}
	cmd := parts[0]

	switch cmd {
	case "home":
		m.state = HomeState
		m.frameCounter = 0
		return m, animationTick()

	case "about":
		m.state = AboutState
		m.bioIndex = 0
		return m, typingTick()

	case "quit", "exit":
		return m, tea.Quit

	default:
		m.errMsg = "Unknown command: " + input + ". Try /home, /about, or /quit."
		return m, nil
	}
}
