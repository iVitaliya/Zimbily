package utils

import (
	"fmt"

	"github.com/iVitaliya/colors-go"
)

type Logger struct{}
type LogState int

const (
	Default LogState = iota
	Info
	Debug
	Warning
	Error
	Fatal
)

func Log() *Logger {
	return &Logger{}
}

func (l *Logger) Error(message ...any) {
	
	msg := String().Replacer("{}{}{} {} zmb [script]", []string{
		colors.BrightBlack("["),
		colors.Red("ERROR"),
		colors.BrightBlack("]"),
		colors.BrightRed(""),
	})
	fmt.Println()
}