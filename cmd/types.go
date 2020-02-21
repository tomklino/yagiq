package yagiq

type yamlType int

const (
  String yamlType = iota
  Integer
  Dictionary
  List
)

type yamlNode struct {
  key string
  valueType yamlType
  stringVal string
  intVal int
  dictionaryVal map[string]*yamlNode
  ListVal []*yamlNode
  LineReference *listNode
}

type scanner interface{
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
}
