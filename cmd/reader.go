package yagiq

import (
  "errors"
  "strings"
  "fmt"
)

var OutOfLinesError = errors.New("Out Of Lines")

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
  return strings.Split(s[indent*2:], ":")[0], nil
}

func MakeTree(listScanner listScanner) (*yamlNode, error) {
  treeParser, err := NewTreeParser(listScanner)
  if err != nil {
    return nil, err
  }
  for {
    err = treeParser.ParseNextLine()
    if err != nil {
      if err != OutOfLinesError {
        return nil, err
      }
      break;
    }
  }

  return treeParser.Root, nil
}

func NewTreeParser(listScanner listScanner) (*TreeParser, error) {
  node := new(yamlNode)
  TreeParser := &TreeParser{
    listScanner: listScanner,
    Root: node,
    currentParent: node,
    parentIndent: -1,
  }
  return TreeParser, nil
}

func (t *TreeParser) connectNode(n *yamlNode) error {
  parent := t.currentParent
  nodeIndent, err := GetLineIndentation(n.LineReference.content)
  if err != nil {
    return err
  }
  switch parent.ValueType {
  case None:
    parent.ValueType = Dictionary
    parent.DictionaryVal = make(map[string]*yamlNode)
    return t.connectNode(n)
  case Dictionary:
    // if lineIsListItem { return fmt.Errorf("invalid yaml....")}
    if nodeIndent != t.parentIndent + 1 {
      return fmt.Errorf("unexpeted indentation when trying to connect node '%s'", n.LineReference.content)
    }
    if parent.DictionaryVal[n.Key] != nil {
      return fmt.Errorf("duplicate key at line '%s'", n.LineReference.content)
    }
    parent.DictionaryVal[n.Key] = n
  // case List:
  default:
    return fmt.Errorf("unexpeted parent type %v when trying to connect node", parent.ValueType)
  }
  return nil
}

func (t *TreeParser) setParent(n *yamlNode) {
  t.currentParent = n
  t.setParentIndent()
}

func (t *TreeParser) setParentIndent() {
  if t.currentParent == t.Root {
    t.parentIndent = -1
  } else {
    // TODO set the line type to also hold its own indent so there will be no
    //      need to recall GetLineIndentation every time and check for errors
    t.parentIndent , _ = GetLineIndentation(t.currentParent.LineReference.content)
  }
}

func (t *TreeParser) ParseNextLine() error {
  if !t.listScanner.Scan() {
    return OutOfLinesError
  }
  l := t.listScanner.Line()

  // TODO smells like a function NewYamlNode(l *listNode)
  n := new(yamlNode)
  n.LineReference = l
  keyName, err := getKeyName(l.content)
  if err != nil {
    return err
  }
  n.Key = keyName

  indent, err := GetLineIndentation(l.content)
  if err != nil {
    return err
  }
  for indent != t.parentIndent + 1 && t.currentParent != t.Root {
    t.currentParent = t.currentParent.ParentNode
    t.setParentIndent()
  }
  if indent != t.parentIndent + 1 {
    return fmt.Errorf("unexpeted indentation %d in line '%s'. expected %d", indent, l.content, t.parentIndent + 1)
  }

  switch {
  case isLineObjectKey(l.content):
    n.ValueType = Dictionary
    n.DictionaryVal = make(map[string]*yamlNode)
    t.connectNode(n)
    t.setParent(n)
  case isLineIntegerKey(l.content):
    n.ValueType = Integer
    n.IntVal = 0 // TODO parse int value
    t.connectNode(n)
  case isLineStringKey(l.content):
    n.ValueType = String
    n.StringVal = parseStringFromLine(l.content)
    t.connectNode(n)
  }
  return nil
}
