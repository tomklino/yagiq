package yagiq

import (
  "testing"
)

func TestIsLineObjectKey(t *testing.T) {
  if isLineObjectKey("something:") != true {
    t.Errorf("Expected 'something:' to be true")
  }
}
