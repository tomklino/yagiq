package yagiq

type yamlType int

const (
  String yamlType = iota
  Integer
  Dictionary
  List
)

type yamlNode struct {
  Key string
  ParentNode *yamlNode
  ValueType yamlType
  StringVal string
  IntVal int
  DictionaryVal map[string]*yamlNode
  ListVal []*yamlNode
  LineReference *listNode
}

type TreeParser struct {
  listScanner
  root        *yamlNode
  current     *yamlNode
}

type listScanner interface {
  Scan() bool
  Line() *listNode
}

type scanner interface {
  Scan() bool
  Text() string
}

type mockScanner struct {
  lines []string
  cursor int
}

type listNode struct {
  content string
  next *listNode
}

type list struct {
  head *listNode
  tail *listNode
}
