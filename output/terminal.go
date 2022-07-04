package output

import (
	"fmt"
	"strings"
)

const NoTextColor = "\033[0m"
const GrayTextColor = "\033[0;37m"
const RedTextColor = "\033[0;31m"
const GreenTextColor = "\033[0;32m"

type TerminalOutput struct {
	NumberOfLines int
	LabelPadding  int
}

func NewTerminalOutput(numberOfLines, longestServerName int) *TerminalOutput {
	createBlankBufferLines(numberOfLines)

	const extraPadding = 20
	return &TerminalOutput{
		numberOfLines,
		longestServerName + extraPadding,
	}
}

func (to TerminalOutput) WriteInfo(label, status string, index int) {
	to.writeToLineAtIndex(label, status, GrayTextColor, index)
}

func (to TerminalOutput) WriteError(label, status string, index int) {
	to.writeToLineAtIndex(label, status, RedTextColor, index)
}

func (to TerminalOutput) WriteSuccess(label, status string, index int) {
	to.writeToLineAtIndex(label, status, GreenTextColor, index)
}

func (to TerminalOutput) writeToLineAtIndex(label, status, color string, index int) {
	lines := to.numberOfLinesToMoveToIndex(index)
	str := FormatLine(label, status, to.LabelPadding)
	fmt.Print(moveUp(lines) + clearLine() + color + str + NoTextColor + moveDown(lines))
}

func createBlankBufferLines(numberOfLines int) {
	fmt.Print(strings.Repeat("\n", numberOfLines))
}

func moveUp(lines int) string {
	return fmt.Sprintf("\x1b[%dF", lines)
}

func moveDown(lines int) string {
	return fmt.Sprintf("\x1b[%dE", lines)
}

func clearLine() string {
	return "\x1b[2K"
}

func (to TerminalOutput) numberOfLinesToMoveToIndex(index int) int {
	return to.NumberOfLines - index
}
