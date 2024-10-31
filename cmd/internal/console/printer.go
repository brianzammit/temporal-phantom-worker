package console

import (
	"github.com/fatih/color"
)

// Success prints a success message in green
func Success(message string, args ...interface{}) {
	green := color.New(color.FgGreen)
	green.Printf("✅  Success: "+message+"\n", args...)
}

// Warn prints a warning message in yellow
func Warn(message string, args ...interface{}) {
	yellow := color.New(color.FgYellow)
	yellow.Printf("⚠️  Warning: "+message+"\n", args...)
}

// Error prints an error message in red
func Error(message string, args ...interface{}) {
	red := color.New(color.FgRed)
	red.Printf("❌  Error: "+message+"\n", args...)
}
