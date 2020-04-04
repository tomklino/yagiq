package yagiq

import (
  "testing"
)

func TestParseStringFromLine(t *testing.T) {
  tests := []struct{
    in string
    out string
  }{
    {"key: \"value\"", "value"},
    {"key: value", "value"},
    {"key: value # comment", "value"},
  }
  for _, tst := range tests {
    if result := parseStringFromLine(tst.in); result != tst.out {
      t.Errorf("exptected '%s' to output '%s', but got '%s'", tst.in, tst.out, result)
    }
  }
}
