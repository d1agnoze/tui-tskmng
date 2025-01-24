package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type binding interface {
	tea.KeyType | string
}

type keyNamespace struct{}

var key = keyNamespace{}

func (keyNamespace) matches(m tea.KeyMsg, s ...any) bool {
	if len(s) == 0 {
		panic("no keys provided")
	}
	for _, key := range s {
		switch t := key.(type) {
		case tea.KeyType, string:
			if m.String() == fmt.Sprintf("%s", t) {
				return true
			}
		default:
			panic("invalid key type")
		}
	}
	return false
}
