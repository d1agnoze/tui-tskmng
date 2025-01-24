package ui

import (
	"github.com/d1agnoze/tui-tskmng/internal/parser"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	table table.Model
	tasks []parser.Task
}

func New(i []parser.Task) model {
	rows := []table.Row{}
	cols := []table.Column{
		{Title: "Name", Width: 15},
		{Title: "Type", Width: 4},
		{Title: "To", Width: 18},
		{Title: "Body", Width: 30},
		{Title: "Subject", Width: 10},
	}

	for _, t := range i {
		rows = append(rows, table.Row{t.Name, t.Type, t.Params.To, t.Params.Body, t.Params.Subject})
	}

	return model{tasks: i, table: newTable(cols, rows)}
}

func newTable(cols []table.Column, rows []table.Row) table.Model {
	opts := []table.Option{
		table.WithRows(rows),
		table.WithColumns(cols),
		table.WithFocused(true),
		table.WithHeight(7),
		table.WithStyles(getTabelStyle()),
	}
	return table.New(opts...)
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if cmd := m.keyBindings(msg); cmd != nil {
			return m, cmd
		}
	}

	newtable, cmd := m.table.Update(msg)
	m.table = newtable
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func (m model) Run() error {
	_, err := tea.NewProgram(m, tea.WithAltScreen()).Run()
	return err
}

func (m *model) keyBindings(msg tea.KeyMsg) tea.Cmd {
	switch {
	case key.matches(msg, tea.KeyCtrlC, "q"):
		return tea.Quit
	}
	return nil
}
