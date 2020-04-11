package yagiq

import (
  "errors"
  "strings"
  "fmt"
)

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

func makeObject(listScanner listScanner) (map[string]*yamlNode, error) {
  result := make(map[string]*yamlNode)
  l := listScanner.Line()
  baseIndent, err := GetLineIndentation(l.content)
  if err != nil {
    return nil, err
  }

  var indent int
  for {
    l := listScanner.Line();

    indent, err = GetLineIndentation(l.content)
    if err != nil {
      return nil, err
    }
    if indent != baseIndent {
      if indent > baseIndent {
        return nil, fmt.Errorf("unexpeted indentation %d, base indent is %d", indent, baseIndent)
      }
      break;
    }

    keyName, err := getKeyName(l.content)
    if err != nil {
      return nil, err
    }
    result[keyName] = new(yamlNode)

    switch {
    case isLineObjectKey(l.content):
      result[keyName].ValueType = Dictionary
      if !isValStringEmtpy(l.content) {
        // TODO isValStringValidJson? if so, parse it as json, if not, return an error
        return nil, fmt.Errorf("object key found in line but object is invalid '%s'", l.content)
      }
      listScanner.Scan()
      object, err := makeObject(listScanner)
      if err != nil {
        return nil, err
      }
      result[keyName].DictionaryVal = object
    case isLineIntegerKey(l.content):
      result[keyName].ValueType = Integer
      // TODO result[keyName].IntVal = <parsed int val>
    case isLineStringKey(l.content):
      result[keyName].ValueType = String
      result[keyName].StringVal = parseStringFromLine(l.content)
    }

    if !listScanner.Scan() {
      break;
    }
  }

  return result, nil
}

func MakeTree(listScanner listScanner) (*yamlNode, error) {
  yamlHead := new(yamlNode)

  if !listScanner.Scan() {
    return nil, errors.New("no lines passed")
  }
  l := listScanner.Line()

  yamlHead.LineReference = l
  yamlHead.ValueType = Dictionary
  object, err := makeObject(listScanner)
  yamlHead.DictionaryVal = object
  if err != nil {
    return nil, err
  }
  return yamlHead, nil
}

func NewTreeParser(listScanner listScanner) (*TreeParser, error) {
  node := new(yamlNode)
  node.DictionaryVal = make(map[string]*yamlNode)
  TreeParser := &TreeParser{
    listScanner: listScanner,
    root: node,
    current: node,
  }
  return TreeParser, nil
}

func (t *TreeParser) ParseNextLine() error {
  if !t.listScanner.Scan() {
    return nil
  }
  l := t.listScanner.Line()

  keyName, err := getKeyName(l.content)
  if err != nil {
    return err
  }

  n := new(yamlNode)
  n.LineReference = l
  switch {
  case isLineObjectKey(l.content):
    n.ValueType = Dictionary
    n.DictionaryVal = make(map[string]*yamlNode)
    n.ParentNode = t.current
    if n.ParentNode.DictionaryVal[keyName] != nil {
      return fmt.Errorf("duplicate key in line '%s'", l.content)
    }
    n.ParentNode.DictionaryVal[keyName] = n
  }
  if t.root == nil {
    t.root = n
  }
  return nil
}
