package main

import (
	"bytes"
	"fmt"
)

// AnsiRenderer exposes functionality using ANSI
type AnsiRenderer struct {
	buffer  *bytes.Buffer
	formats *ansiFormats
}

func (r *AnsiRenderer) carriageForward() {
	r.buffer.WriteString(fmt.Sprintf(r.formats.left, 1000))
}

func (r *AnsiRenderer) setCursorForRightWrite(text string, offset int) {
	strippedLen := r.formats.lenWithoutANSI(text) + -offset
	r.buffer.WriteString(fmt.Sprintf(r.formats.right, strippedLen))
}

func (r *AnsiRenderer) changeLine(numberOfLines int) {
	position := "B"
	if numberOfLines < 0 {
		position = "F"
		numberOfLines = -numberOfLines
	}
	r.buffer.WriteString(fmt.Sprintf(r.formats.linechange, numberOfLines, position))
}

func (r *AnsiRenderer) creset() {
	r.buffer.WriteString(r.formats.creset)
}

func (r *AnsiRenderer) print(text string) {
	r.buffer.WriteString(text)
	r.clearEOL()
}

func (r *AnsiRenderer) clearEOL() {
	r.buffer.WriteString(r.formats.clearOEL)
}

func (r *AnsiRenderer) string() string {
	return r.buffer.String()
}

func (r *AnsiRenderer) saveCursorPosition() {
	r.buffer.WriteString(r.formats.saveCursorPosition)
}

func (r *AnsiRenderer) restoreCursorPosition() {
	r.buffer.WriteString(r.formats.restoreCursorPosition)
}