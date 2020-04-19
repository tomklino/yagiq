package yagiq

import (
  "testing"
)

var dummyNode = &yamlNode{
  Key: "obj",
  ValueType: Dictionary,
  LineReference: &listNode{
    content: "obj:",
  },
}

func TestNewTreeParser(t *testing.T) {
  mockScanner := CreateMockScanner([]string{})
  dummyScanner := NewFListScanner(mockScanner)

  treeParser, err := NewTreeParser(dummyScanner)
  if err != nil {
    t.Errorf("did not expect an error while creating a new treeParser: %s", err)
  }
  if treeParser.Root == nil {
    t.Errorf("expected tree root to be initialized")
  }
  if treeParser.parentIndent != -1 {
    t.Errorf("expected the parentIndent variable to be initialized to -1")
  }
}

func TestConnectNode(t *testing.T) {
  mockScanner := CreateMockScanner([]string{})
  dummyScanner := NewFListScanner(mockScanner)

  tree, _ := NewTreeParser(dummyScanner)

  tree.connectNode(dummyNode)

  if tree.Root.ValueType != Dictionary {
    t.Errorf("expected the root of the tree to become a dictionary after connecting a non-list node to it")
  }
  if tree.Root.DictionaryVal["obj"] == nil {
    t.Errorf("expected root to have a child under the key 'obj'")
  }
  if tree.Root.DictionaryVal["obj"] != dummyNode {
    t.Errorf("expected the node at root.obj to equal the dummy yaml node")
  }
}

func TestSetParentIndent(t *testing.T) {
  mockScanner := CreateMockScanner([]string{})
  dummyScanner := NewFListScanner(mockScanner)

  tree, _ := NewTreeParser(dummyScanner)

  tree.connectNode(dummyNode)
  tree.setParentIndent()
  if tree.parentIndent != 0 {
    t.Errorf("expected the parent indent to be set to 0 after connecting the dummy node to it")
  }
}
