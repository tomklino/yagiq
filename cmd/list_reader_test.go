package yagiq

import (
  "testing"
)

func TestListScanner(t *testing.T) {
  mockScanner := CreateMockScanner(dummyLines)
  listScanner := NewFListScanner(mockScanner)
  var firstLine, secondLine *listNode
  for i, line := range dummyLines {
    listScanner.Scan()
    l := listScanner.Line()
    if l.content != line {
      t.Errorf("exptected to read '%s', got '%s'", line, l.content)
    }
    if i == 0 {
      firstLine = l
    }
    if i == 1 {
      secondLine = l
    }
  }
  if ok := listScanner.Scan(); ok != false {
    t.Errorf("expected ok to be false when reading after the last line, but got true")
  }
  if firstLine.next != secondLine {
    t.Errorf("expected the first line to point to the second line")
  }
}
