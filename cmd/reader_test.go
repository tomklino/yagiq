package yagiq

import (
  "testing"
)

type scanner struct {
  lines []string
  cursor int
}

func min(a, b int) int {
  if (a > b) {
    return b
  }
  return a
}

func CreateMockScanner(lines []string) *scanner {
  scanner := &scanner{
    lines: lines,
    cursor: -1,
  }
  return scanner
}

func (s *scanner) Scan() {
  s.cursor = min(s.cursor + 1, len(s.lines))
}

func (s *scanner) Text() string {
  return s.lines[s.cursor]
}

func TestReadToList(t *testing.T) {
  mockScanner := CreateMockScanner([]string{
    "object:",
    "  key: \"value\"",
    "  another: \"val2\"",
  })
  lineList := ReadToList(mockScanner)

  if firstLine := lineList.head.content; firstLine != "object:" {
    t.Errorf("first line is %s; want \"object:\"", firstLine)
  }
}
