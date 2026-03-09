package ui

import "github.com/charmbracelet/lipgloss"

func (m Model) viewHome() string {
	if len(m.frames) == 0 {
		return "No frames loaded."
	}

	frame := m.frames[m.frameCounter%len(m.frames)]

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(false)

	rendered := style.Render(frame)

	return lipgloss.Place(
		m.terminalW,
		m.terminalH-4,
		lipgloss.Center,
		lipgloss.Center,
		rendered,
	)
}
