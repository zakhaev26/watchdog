package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type model struct {
	focusIndex            int
	inputs                []textinput.Model
	cursorMode            cursor.Mode
	ES_HOST               string
	ES_API_KEY            string
	CRITICAL_LOG_NODE_ID  string
	FREQUENT_LOG_NODE_ID  string
	MAILING_ID            string
	EMAIL_SENDER_NAME     string
	EMAIL_SENDER_ADDRESS  string
	EMAIL_SENDER_PASSWORD string
}

var k model

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 8),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Enter ElasticSearch Host URI"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.CharLimit = 190
		case 1:
			t.Placeholder = "Enter Kibana API Key"
			t.CharLimit = 64
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'

		case 2:
			t.Placeholder = "Enter the CRITICAL LOG name with which kafka Topics and ElasticSearch Indices will be created"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle

		case 3:
			t.Placeholder = "Enter the FREQUENT LOG name with which kafka Topics and ElasticSearch Indices will be created"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle

		case 4:
			t.Placeholder = "Enter the Mailing ID where Critical Alerts will be sent"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 5:
			t.Placeholder = "Enter the SMTP Email Sender Name (Gmail)"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 6:
			t.Placeholder = "Enter the SMTP Email Address (Gmail)"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 7:
			t.Placeholder = "Enter the SMTP Email Sender Password (Gmail)"
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				cmds[i] = m.inputs[i].Cursor.SetMode(m.cursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				k.ES_HOST = m.inputs[0].Value()
				k.ES_API_KEY = m.inputs[1].Value()
				k.CRITICAL_LOG_NODE_ID = m.inputs[2].Value()
				k.FREQUENT_LOG_NODE_ID = m.inputs[3].Value()
				k.MAILING_ID = m.inputs[4].Value()
				k.EMAIL_SENDER_NAME = m.inputs[5].Value()
				k.EMAIL_SENDER_ADDRESS = m.inputs[6].Value()
				k.EMAIL_SENDER_PASSWORD = m.inputs[7].Value()

				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}

func main() {
	m := initialModel()
	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}

	// Access the values from the model after the program has exited
	mail_service_env := `EMAIL_SENDER_NAME=` + k.EMAIL_SENDER_NAME +
		`
EMAIL_SENDER_ADDRESS=` + k.EMAIL_SENDER_ADDRESS +
		`
EMAIL_SENDER_PASSWORD=` + k.EMAIL_SENDER_PASSWORD

	frequent_producer_env := `FREQUENT_LOG_NODE_ID=` + k.FREQUENT_LOG_NODE_ID
	critical_producer_env := `CRITICAL_LOG_NODE_ID=` + k.CRITICAL_LOG_NODE_ID

	frequent_consumer_env := `FREQUENT_LOG_NODE_ID=` + k.FREQUENT_LOG_NODE_ID +
		`
ES_HOST=` + k.ES_HOST +
		`
ES_API_KEY=` + k.ES_API_KEY

	critical_consumer_env := `CRITICAL_LOG_NODE_ID=` + k.CRITICAL_LOG_NODE_ID +
		`
ES_HOST=` + k.ES_HOST +
		`
ES_API_KEY=` + k.ES_API_KEY

	err := os.WriteFile("../mail_service/.env", []byte(mail_service_env), 0755)
	if err != nil {
		fmt.Println("Error creating Bash script:", err)
	}

	err = os.WriteFile("../frequent_producer/.env", []byte(frequent_producer_env), 0755)
	if err != nil {
		fmt.Println("Error creating Bash script:", err)
	}

	err = os.WriteFile("../critical_producer/.env", []byte(critical_producer_env), 0755)
	if err != nil {
		fmt.Println("Error creating Bash script:", err)
	}

	err = os.WriteFile("../frequent_consumer/_elastic/.env", []byte(frequent_consumer_env), 0755)
	if err != nil {
		fmt.Println("Error creating Bash script:", err)
	}

	err = os.WriteFile("../critical_consumer/_elastic/.env", []byte(critical_consumer_env), 0755)
	if err != nil {
		fmt.Println("Error creating Bash script:", err)
	}
}
