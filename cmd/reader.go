package yagiq

import (
  "errors"
  "strings"
  "fmt"
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

func getKeyName(s string) (string, error) {
  indent, err := GetLineIndentation(s)
  if err != nil {
    return "", err
  }
  return strings.Split(s[indent+1:], ":")[0], nil
}

func makeObject(l *listNode) (map[string]*yamlNode, *listNode, error) {
  result := make(map[string]*yamlNode)
  baseIndent, err := GetLineIndentation(l.content)
  if err != nil {
    return nil, nil, err
  }

  var indent int
  for l != nil {
    indent, err = GetLineIndentation(l.content)
    if err != nil {
      return nil, nil, err
    }
    if (indent != baseIndent) {
      break;
    }

    keyName, err := getKeyName(l.content)
    if err != nil {
      return nil, nil, err
    }
    result[keyName] = new(yamlNode)

    result[keyName].ValueType = String // TODO implement this
    result[keyName].StringVal = "value" // TODO implement this
    l = l.next
  }

  if(indent > baseIndent) {
    return nil, nil, fmt.Errorf("unexpeted indentation %d", indent)
  }
  return result, l, nil
}

func MakeTree(l *listNode) (*yamlNode, error) {
  result := new(yamlNode)

  if(isLineObjectKey(l.content)) {
    result.Key = "object" // TODO implement this
    result.ValueType = Dictionary
    result.LineReference = l
    object, l, err := makeObject(l.next)
    if err != nil {
      return nil, err
    }
    result.DictionaryVal = object
    if l == nil || l.next == nil {
      return result, nil
    }
  }
  return result, nil
}
