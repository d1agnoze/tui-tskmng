package parser

import "fmt"

type ast struct {
	Tasks []*Task `parser:"@@*"`
}

type Task struct {
	Type   string      `parser:"@\"task\""`
	Name   string      `parser:"@String"`
	Params *TaskParams `parser:"'{' @@ '}'"`
}

type TaskParams struct {
	To      string `parser:"'to' '=' @String"`
	Subject string `parser:"'subject' '=' @String"`
	Body    string `parser:"'body' '=' @String"`
}

func (i *ast) String() string {
	out := ""
	for _, v := range i.Tasks {
		out += fmt.Sprintf("%s\n", v)
	}
	return out
}

func (s *Task) String() string {
	return fmt.Sprintf("Task: %s | %s -> %s: %s ", s.Name, s.Params.To, s.Params.Subject, s.Params.Body)
}

/*
-------------------
Base value implementation
-------------------
*/

type Value interface{ value() }

type Float struct {
	Value float64 `parser:"@Float"`
}

type Int struct {
	Value int `parser:"@Int"`
}

type String struct {
	Value string `parser:"@String"`
}

type Bool struct {
	Value bool `parser:"@('true' | 'false')"`
}

func (Float) value()  {}
func (Int) value()    {}
func (String) value() {}
func (Bool) value()   {}
