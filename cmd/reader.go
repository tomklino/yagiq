package yagiq

func ReadToList(s scanner) *list {
  list := &list{}
  tracer := &list.head
  for s.Scan() {
    currentNode := &listNode{
      content: s.Text(),
    }
    *tracer = currentNode
    tracer = &currentNode.next
  }
  return list
}
