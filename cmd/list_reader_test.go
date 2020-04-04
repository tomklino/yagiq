package yagiq

import (
  "testing"
)

func TestReadNext(t *testing.T) {
  mockScanner := CreateMockScanner(dummyLines)
  listReader := &FListReader{mockScanner}
  var firstLine, secondLine *listNode
  for i, line := range dummyLines {
    res, _ := listReader.ReadNext()
    if res.content != line {
      t.Errorf("exptected to read '%s', got '%s'", line, res.content)
    }
    if i == 0 {
      firstLine = res
    }
    if i == 1 {
      secondLine = res
    }
  }
  if _, ok := listReader.ReadNext(); ok != false {
    t.Errorf("expected ok to be false when reading after the last line, but got true")
  }
  if firstLine.next != secondLine {
    t.Errorf("expected the first line to point to the second line")
  }
}
