package config

type Window struct {
	Panes  []Pane
	Layout string
	Sync   *bool
}
