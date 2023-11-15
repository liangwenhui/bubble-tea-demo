package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

const (
	lineLimit = 30
)

type ChatModel struct {

	//chat window
	view     viewport.Model
	textArea textarea.Model
	messages []string
}

var _ tea.Model = &ChatModel{}

var senderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))

func NewChatRoom() *ChatModel {

	ta := textarea.New()
	ta.Placeholder = "say something... enter to send message"
	ta.Focus()

	ta.Prompt = "|🖊 "
	ta.CharLimit = 250
	ta.SetHeight(3)
	ta.SetWidth(lineLimit)

	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.ShowLineNumbers = false

	vp := viewport.New(lineLimit+6, 5)
	vp.SetContent("welcome to chat room. 😀")
	ta.KeyMap.InsertNewline.SetEnabled(false)

	return &ChatModel{
		textArea: ta,
		view:     vp,
	}
}

func (c *ChatModel) Init() tea.Cmd {
	return textarea.Blink
}

func (c *ChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var vc, tc tea.Cmd

	c.view, vc = c.view.Update(msg)
	c.textArea, tc = c.textArea.Update(msg)
	switch t := msg.(type) {
	case tea.KeyMsg:
		switch t.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return c, tea.Quit
		case tea.KeyEnter:
			input := c.textArea.Value()
			c.messages = append(c.messages, senderStyle.Render("Your:")+splitMessage(input))
			c.view.SetContent(strings.Join(c.messages, "\n"))
			c.textArea.Reset()
			c.view.GotoBottom()
		}
	}
	return c, tea.Batch(vc, tc)
}

func splitMessage(input string) string {
	// only en
	if len(input) <= lineLimit {
		return input
	}
	s := input[:lineLimit+1]
	b := input[lineLimit+1:]

	return s + "\n      " + splitMessage(b)

}

func (c *ChatModel) View() string {
	return fmt.Sprintf(
		"%s\n%s",
		c.view.View(),
		c.textArea.View(),
	) + "\n"
}

func main() {
	model := NewChatRoom()
	program := tea.NewProgram(model)
	if _, err := program.Run(); err != nil {
		panic(err)
	}
}
