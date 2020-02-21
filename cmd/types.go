package yagiq

type scanner interface{
  Scan() bool
  Text() string
}

type mockScanner struct {
  lines []string
  cursor int
}

type listNode struct {
  content string
  next *listNode
}

type list struct {
  head *listNode
}
