package yagiq

import (
  "testing"
)

var dummyLines = []string{
  "object:",
  "  key: \"value\"",
  "  another: \"val2\"",
}

func CreateMockScanner(lines []string) *mockScanner {
  scanner := &mockScanner{
    lines: lines,
    cursor: -1,
  }
  return scanner
}

func (s *mockScanner) Scan() bool {
  s.cursor = s.cursor + 1
  if s.cursor >= len(s.lines) {
    return false
  }
  return true
}

func (s *mockScanner) Text() string {
  return s.lines[s.cursor]
}

func TestReadToList(t *testing.T) {
  mockScanner := CreateMockScanner(dummyLines)
  lineList := ReadToList(mockScanner)

  firstLine := lineList.head
  if firstLine.content != dummyLines[0] {
    t.Errorf("first line is \"%s\"; want \"%s\"", firstLine.content, dummyLines[0])
  }
  secondLine := firstLine.next
  if secondLine.content != dummyLines[1] {
    t.Errorf("second line is \"%s\"; want \"%s\"", secondLine.content, dummyLines[1])
  }
}

// func TestGetLineIndentation(t *testing.T) {
//
// }
