package yagiq

import (
  "strings"
)

func parseStringFromLine(s string) string {
  colonIndex := strings.IndexRune(s, ':')
  valuePartString := strings.Trim(s[colonIndex+1:], "\t ")
  if qs := strings.IndexRune(valuePartString, '"'); qs != -1 {
    if qe := strings.IndexRune(valuePartString[qs+1:], '"'); qe != -1 {
      // 2 quote symbols found, returning the value between the quotes
      return valuePartString[qs+1:qe+1]
    }
  }
  // no quotes, trimming out the comment if necessary and returning
  if commentIndex := strings.IndexRune(valuePartString, '#'); commentIndex != -1 {
    valuePartString = valuePartString[:commentIndex]
  }
  return strings.Trim(valuePartString, "\t ")
}
