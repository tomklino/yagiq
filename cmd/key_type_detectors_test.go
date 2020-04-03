package yagiq

import (
  "testing"
)

func TestIsLineObjectKey(t *testing.T) {
  if isLineObjectKey("something:") != true {
    t.Errorf("Expected 'something:' to be true")
  }
}

func TestIsLineIntegerKey(t *testing.T) {
  var integerTests = []struct{
    in string
    out bool
  }{
    {"key: value", false},
    {"key: \"value\"", false},
    {"key:", false},
    {"key:5", true},
    {"key: 5", true},
  }
  for _, test := range integerTests {
    if isLineIntegerKey(test.in) != test.out {
      t.Errorf("the line '%s' is exptected to be %v, but got %v", test.in, test.out, !test.out)
    }
  }
}

func TestIsLineStringKey(t *testing.T) {
  var stringTests = []struct{
    in string
    out bool
  }{
    {"key: value", true},
    {"key: \"value\"", true},
    {"key:", false},
    {"key: 5", false},
  }
  for _, test := range stringTests {
    if isLineStringKey(test.in) != test.out {
      t.Errorf("the line '%s' is exptected to be %v, but got %v", test.in, test.out, !test.out)
    }
  }
}
