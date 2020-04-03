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
  if isLineStringKey("key: value") != true {
    t.Errorf("the line 'key: value' is a string key, but got false")
  }
  if isLineStringKey("key: \"chooo\"") != true {
    t.Errorf("the line 'key: \"chooo\"' is a string key, but got false")
  }
  if isLineStringKey("key:") != false {
    t.Errorf("the line 'key:' is not a string key, but got true")
  }
  if isLineStringKey("key: 5") != false {
    t.Errorf("the line 'key: 5' is not a string key, but got true")
  }
}
