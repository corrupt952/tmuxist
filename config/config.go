package config

type Config struct {
	Name    string
	Root    string
	Attach  *bool
	Windows []Window
}
