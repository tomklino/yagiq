package yagiq

import (
  "testing"
)

type detectorTestCase struct{
  in string
  out bool
}

func TestIsLineObjectKey(t *testing.T) {
  var objectTests = []detectorTestCase{
    {"something:", true},
  }
  for _, test := range objectTests {
    if isLineObjectKey(test.in) != test.out {
      t.Errorf("line '%s' expected to be %v, but got %v", test.in, test.out, !test.out)
    }
  }
}

func TestIsLineIntegerKey(t *testing.T) {
  var integerTests = []detectorTestCase{
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
  var stringTests = []detectorTestCase{
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
