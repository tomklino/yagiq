package yagiq

type FListReader struct {
  scanner
}

func (f *FListReader) ReadNext() (*listNode, bool) {
  // NOTE may need to add tracer or tailNode in closure to be able to link
  //      last node to next node
  if !f.scanner.Scan() {
    return nil, false
  }
  line := f.scanner.Text()
  result := &listNode{
    content: line,
  }
  return result, true
}
