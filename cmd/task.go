package cmd

// CurrentTask holds the struct of the currently
// highlighted Task.
var CurrentTask Task

// Task model.
type Task struct {
	Name        string
	Description string
	Logs        string
}
