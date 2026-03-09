package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/0xf6i/cli/internal/ui/components"
)

type SessionState int

const (
	HomeState SessionState = iota
	AboutState
	EducationState
	ExperienceState
)

type ExperienceSubView int

const (
	ExperienceOverview ExperienceSubView = iota
	ExperienceWork
	ExperienceAcademia
)

type Model struct {
	textInput textinput.Model
	state     SessionState
	terminalH int
	terminalW int

	// Home view — animation
	frameCounter int
	frames       []string

	// About view — streaming text
	bioText  string
	bioIndex int

	// Experience view
	expSubView ExperienceSubView

	// Error feedback
	errMsg string
}

func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Try /home, /about, or whatever you feel like big dog."
	ti.Focus()
	ti.ShowSuggestions = true
	ti.SetSuggestions([]string{"/home", "/about", "/education", "/experience", "/quit"})

	frames, err := components.LoadFrames("assets/frames")
	if err != nil {
		frames = []string{"[ error loading frames ]"}
	}

	return Model{
		state:     HomeState,
		textInput: ti,
		frames:    frames,
		bioText:   bioContent,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, animationTick())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.terminalH = msg.Height
		m.terminalW = msg.Width
		return m, nil

	case animationTickMsg:
		if m.state == HomeState {
			m.frameCounter++
			return m, animationTick()
		}
		return m, nil

	case typingTickMsg:
		if m.state == AboutState && m.bioIndex < len(m.bioText) {
			m.bioIndex++
			return m, typingTick()
		}
		return m, nil

	case tea.KeyMsg:
		m.errMsg = ""
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			input := strings.TrimSpace(m.textInput.Value())
			m.textInput.SetValue("")
			if input != "" {
				m, cmd = m.handleCommand(input)
				m.textInput, _ = m.textInput.Update(msg)
				return m, cmd
			}
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.terminalW == 0 {
		return "Loading..."
	}

	var content string
	switch m.state {
	case HomeState:
		content = m.viewHome()
	case AboutState:
		content = m.viewAbout()
	default:
		content = m.viewHome()
	}

	// Error line
	errLine := ""
	if m.errMsg != "" {
		errLine = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(m.errMsg) + "\n"
	}

	input := m.textInput.View()

	return content + "\n\n" + errLine + input
}
