package yagiq

func isLineObjectKey(s string) bool {
  // TODO trim comments
  return s[len(s)-1] == ':'
}
