package yagiq

import (
  "testing"
)

var dummyLines = []string{
  "object:",
  "  key: \"value\"",
  "  another: \"val2\"",
}

func CreateMockScanner(lines []string) *mockScanner {
  scanner := &mockScanner{
    lines: lines,
    cursor: -1,
  }
  return scanner
}

func (s *mockScanner) Scan() bool {
  s.cursor = s.cursor + 1
  if s.cursor >= len(s.lines) {
    return false
  }
  return true
}

func (s *mockScanner) Text() string {
  return s.lines[s.cursor]
}

func TestReadToList(t *testing.T) {
  mockScanner := CreateMockScanner(dummyLines)
  lineList := ReadToList(mockScanner)

  firstLine := lineList.head
  if firstLine.content != dummyLines[0] {
    t.Errorf("first line is \"%s\"; want \"%s\"", firstLine.content, dummyLines[0])
  }
  secondLine := firstLine.next
  if secondLine.content != dummyLines[1] {
    t.Errorf("second line is \"%s\"; want \"%s\"", secondLine.content, dummyLines[1])
  }
}

func TestGetLineIndentation(t *testing.T) {
  indent, err := GetLineIndentation(dummyLines[0]) //  "object:" (0)
  if err != nil {
    t.Errorf("indent for %s was not successful; returned with unexpeted error %s", dummyLines[0], err)
  }
  if indent != 0 {
    t.Errorf("indent is %d; want 0", indent)
  }

  indent, err =  GetLineIndentation(dummyLines[1]) // "  key: \"value\"" (1)
  if err != nil {
    t.Errorf("indent for %s was not successful; returned with unexpeted error %s", dummyLines[1], err)
  }
  if indent != 1 {
    t.Errorf("indent is %d; want 1", indent)
  }

  indent, err = GetLineIndentation(" invalid") // should return error
  if err == nil {
    t.Errorf("expected to fail with 'illegal number of spaces for indentation' but error is nil")
  }
}

func TestGetKeyName(t *testing.T) {
  result, err := getKeyName("  key:")
  if err != nil {
    t.Errorf("got error while trying to get key name: %s", err)
  }
  if result != "key" {
    t.Errorf("expected 'key', got '%s'", result)
  }
}

func TestMakeTree(t *testing.T) {
  mockScanner := CreateMockScanner(dummyLines)
  dummyList := ReadToList(mockScanner)
  dummyTree, err := MakeTree(dummyList.head)
  if err != nil {
    t.Errorf("make tree returned an unexpeted error %s", err)
    return
  }
  if dummyTree.Key != "object" {
    t.Errorf("head of tree keyname is %s; want 'object'", dummyTree.Key)
  }
  if dummyTree.ValueType != Dictionary {
    t.Errorf("head of tree type is %v; want yamlType.Dictionary", dummyTree.ValueType)
  }
  if dummyTree.LineReference != dummyList.head {
    t.Errorf("head of tree is not referencing the first line in the list")
  }
  if dummyTree.DictionaryVal["key"].ValueType != String {
    t.Errorf("exptected the type of 'object.key' to be string")
  }
  if dummyTree.DictionaryVal["another"].ValueType != String {
    t.Errorf("exptected the type of 'object.another' to be string")
  }
  if dummyTree.DictionaryVal["another"].StringVal != "val2" {
    t.Errorf("exptected the value to 'objecty.another' to be 'val2' but got '%s'", dummyTree.DictionaryVal["another"].StringVal)
  }
}
