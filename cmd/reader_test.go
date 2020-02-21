package yagiq

type scanner struct {
  lines []string
  cursor int
}

func min(a, b int) {
  if (a > b) {
    return b
  }
  return a
}

func (s *scanner) Scan() {
  if s.cursor == nil {
    s.cursor = 0
  } else {
    s.cursor = min(s.cursor + 1, len(s.lines))
  }
}

func (s *scanner) Text() {
  return lines[s.cursor]
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
