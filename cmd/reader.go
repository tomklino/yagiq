package yagiq

import (
  "errors"
  "strings"
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
