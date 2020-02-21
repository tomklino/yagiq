package yagiq

import (
  "errors"
)

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

func GetLineIndentation(s string) (int, error) {
  indents := 0
  for i := 0; i < len(s); i += 2 {
    if s[i] == ' ' {
      if s[i+1] != ' ' {
        return 0, errors.New("illegal number of spaces for indentation")
      }
      indents++
    } else {
      break;
    }
  }
  return indents, nil
}
