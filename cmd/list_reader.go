package yagiq

type FListScanner struct {
  scanner
  list *list
  tracer **listNode
}

func NewFListScanner(s scanner) *FListScanner {
  list := &list{}
  tracer := &list.head
  return &FListScanner{s, list, tracer}
}

func (f *FListScanner) Scan() bool {
  if !f.scanner.Scan() {
    return false
  }
  line := f.scanner.Text()
  result := &listNode{
    content: line,
  }
  *f.tracer = result
  f.tracer = &result.next
  f.list.tail = result
  return true
}

func (f *FListScanner) Line() *listNode {
  return f.list.tail
}
