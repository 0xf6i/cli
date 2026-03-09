package ui

import "github.com/charmbracelet/lipgloss"

func (m Model) viewAbout() string {
	visible := m.bioText[:m.bioIndex]

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(m.cfg.Style.AboutColor)).
		Padding(1, 2)

	rendered := style.Render(visible)

	return lipgloss.Place(
		m.terminalW,
		m.terminalH-4,
		lipgloss.Center,
		lipgloss.Center,
		rendered,
	)
}
