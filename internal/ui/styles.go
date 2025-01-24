package ui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	baseStyle          = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
	borderStyle        = lipgloss.NormalBorder()
	BorderForeground   = lipgloss.Color("240")
	selectedForeground = lipgloss.Color("229")
	selectedBackground = lipgloss.Color("57")
)

func getTabelStyle() table.Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.BorderStyle(borderStyle).BorderForeground(BorderForeground).BorderBottom(true).Bold(false)
	s.Selected = s.Selected.Foreground(selectedForeground).Background(selectedBackground).Bold(false)
	return s
}
