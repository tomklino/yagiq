package yagiq

import (
  "strings"
)

func isLineObjectKey(s string) bool {
  valString := strings.Trim(strings.Split(s, ":")[1], "\t ")
  if len(valString) == 0 || valString == "{}" {
    return true
  }
  return false
}

func isLineIntegerKey(s string) bool {
  valString := strings.Trim(strings.Split(s, ":")[1], "\t ")
  if len(valString) == 0 {
    return false
  }
  switch valString[0] {
  case '0','1','2','3','4','5','6','7','8','9':
    return true
  }
  return false
}

func isLineStringKey(s string) bool {
  valString := strings.Trim(strings.Split(s, ":")[1], "\t ")
  if len(valString) == 0 {
    return false
  }
  switch valString[0] {
  case '0','1','2','3','4','5','6','7','8','9':
    return false
  }
  return true
}
