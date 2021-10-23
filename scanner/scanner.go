package scanner

import (
	"bufio"
	"io"
	"strings"
)

const (
	splitStr = " "
)

type commandScanner struct {
	scanner  *bufio.Scanner
	Commands chan []string
}

func NewScanner(r io.Reader) *commandScanner {
	return &commandScanner{
		scanner:  bufio.NewScanner(r),
		Commands: make(chan []string),
	}
}

func (s *commandScanner) ScanLines() {
	defer close(s.Commands)

	for s.scanner.Scan() {
		line := s.scanner.Text()
		s.Commands <- strings.Split(line, splitStr)
	}
}
