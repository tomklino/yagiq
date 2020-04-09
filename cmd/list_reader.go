package yagiq

type FListReader struct {
  scanner
  list *list
  tracer **listNode
}

func NewFListReader(s scanner) *FListReader {
  list := &list{}
  tracer := &list.head
  return &FListReader{s, list, tracer}
}

func (f *FListReader) ReadNext() (*listNode, bool) {
  if !f.scanner.Scan() {
    return nil, false
  }
  line := f.scanner.Text()
  result := &listNode{
    content: line,
  }
  *f.tracer = result
  f.tracer = &result.next
  return result, true
}
