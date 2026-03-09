package ui

import "github.com/charmbracelet/lipgloss"

const bioContent = `Hey, I'm Ludwig — a developer who builds things for the terminal and the web.

I love Go, distributed systems, and developer tooling. When I'm not writing code,
you'll find me exploring new tech, reading about systems design, or tinkering with
side projects like this SSH portfolio you're looking at right now.

Type /home to go back, or /quit to disconnect.`

func (m Model) viewAbout() string {
	visible := m.bioText[:m.bioIndex]

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")).
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
