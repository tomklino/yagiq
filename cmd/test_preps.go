package yagiq

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
