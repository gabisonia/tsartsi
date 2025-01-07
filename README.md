# Tsartsi - ცარცი - Chalk 🖍️

Tsartsi (Georgian: **ცარცი**) means chalk, it is a lightweight Go package for styling terminal output with ANSI escape codes. It allows you to add colors, background colors, and text styles like bold or underline, making your console output visually appealing.

## Purpose
This project is part of my journey to learn and practice Go. While it may not be a production-ready or professional-grade package, it serves as a hands-on way to explore Go's features, modules, and test writing.

Feel free to use it for learning purposes or experiment with it in your own projects!

## Features
🎨 Foreground Colors: Red, Green, Blue, etc.

🖌️ Background Colors: White, Black, Yellow, etc.

✨ Text Styles: Bold, Underline, Invert, etc.

## Installation

To install the Tsartsi package, run:

```bash
go get github.com/gabisonia/tsartsi
```

## Usage

Here’s how to use Tsartsi in your Go projects:

```go
package main

import (
	"github.com/gabisonia/tsartsi"
)

func main() {
	style := tsartsi.NewStyle(tsartsi.Red, tsartsi.BgWhite, tsartsi.Bold)
	style.Println("Hello, Tsartsi!")
}
```

## Testing
This package includes a test suite to ensure everything works as expected.

To run the tests, use:

```bash
go test ./...

```

## License

Tsartsi is fully free to use, modify, and distribute for any purpose, personal or commercial. 

No attribution is required, but contributions and feedback are always welcome!

