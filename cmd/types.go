package yagiq

type yamlType int

const (
  None yamlType = iota
  String
  Integer
  Dictionary
  List
)

type yamlPath []string

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
  Root          *yamlNode
  currentParent *yamlNode
  parentIndent  int
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
