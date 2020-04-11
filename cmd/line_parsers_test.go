package yagiq

import (
  "testing"
)

func TestIsValStringEmptry(t *testing.T) {
  tests := []struct{
    in string
    out bool
  }{
    {"object:", true},
    {"object:   ", true},
    {"object: {}", false},
    {"object:   # comment", true},
    {"object: string # comment", false},
  }
  for _, tst := range tests {
    if result := isValStringEmtpy(tst.in); result != tst.out {
      t.Errorf("exptected '%s' to output '%v', got '%v'", tst.in, tst.out, result)
    }
  }
}

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
