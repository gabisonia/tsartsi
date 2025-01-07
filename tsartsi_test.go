package tsartsi

import (
	"testing"
)

func TestRender(t *testing.T) {
	style := NewStyle(Red, BgBlack, Bold)
	result := style.Render("Hello, Chalk!")
	expected := "\033[31m\033[40m\033[1mHello, Chalk!\033[0m"
	if result != expected {
		t.Errorf("Render() = %q; want %q", result, expected)
	}
}

func TestRenderNoBackground(t *testing.T) {
	style := NewStyle(Green, "", Underline)
	result := style.Render("Underline Green Text")
	expected := "\033[32m\033[4mUnderline Green Text\033[0m"
	if result != expected {
		t.Errorf("Render() = %q; want %q", result, expected)
	}
}

func TestRenderNoForeground(t *testing.T) {
	style := NewStyle("", BgYellow, Invert)
	result := style.Render("Inverted Yellow Background")
	expected := "\033[43m\033[7mInverted Yellow Background\033[0m"
	if result != expected {
		t.Errorf("Render() = %q; want %q", result, expected)
	}
}

func TestRenderMultipleOptions(t *testing.T) {
	style := NewStyle(Cyan, BgWhite, Bold, Underline)
	result := style.Render("Cyan Bold Underlined")
	expected := "\033[36m\033[47m\033[1m\033[4mCyan Bold Underlined\033[0m"
	if result != expected {
		t.Errorf("Render() = %q; want %q", result, expected)
	}
}

func TestRenderEmptyStyle(t *testing.T) {
	style := NewStyle("", "", "")
	result := style.Render("Plain Text")
	expected := "Plain Text\033[0m"
	if result != expected {
		t.Errorf("Render() = %q; want %q", result, expected)
	}
}

func TestPrint(t *testing.T) {
	style := NewStyle(Blue, BgMagenta, Bold)
	// Visually verify this in the terminal; this test won't fail automatically
	t.Log("Run `go test -v` and visually verify that the following line is styled:")
	style.Print("This is styled text\n")
}

func TestPrintln(t *testing.T) {
	style := NewStyle(White, BgGreen, Underline)
	// Visually verify this in the terminal; this test won't fail automatically
	t.Log("Run `go test -v` and visually verify that the following line is styled:")
	style.Println("This is styled text with a newline")
}
