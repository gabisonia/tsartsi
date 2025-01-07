package tsartsi

import "fmt"

// Style represents a combination of ANSI codes for styling text
type Style struct {
	foreground string
	background string
	options    []string
}

// NewStyle creates a new Style with specified foreground, background, and options
func NewStyle(fg, bg string, options ...string) Style {
	return Style{
		foreground: fg,
		background: bg,
		options:    options,
	}
}

// Render applies the style to the given text and returns the styled string
func (s Style) Render(text string) string {
	codes := s.foreground + s.background
	for _, option := range s.options {
		codes += option
	}
	reset := "\033[0m"
	return codes + text + reset
}

// Convenience functions for common foreground colors
const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

// Convenience functions for common background colors
const (
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

// Text style options
const (
	Bold      = "\033[1m"
	Dim       = "\033[2m"
	Underline = "\033[4m"
	Invert    = "\033[7m"
)

// Print applies the style and prints the styled text
func (s Style) Print(text string) {
	fmt.Print(s.Render(text))
}

// Println applies the style and prints the styled text with a newline
func (s Style) Println(text string) {
	fmt.Println(s.Render(text))
}
