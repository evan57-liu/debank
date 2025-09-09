package ptrutils

// IntPtrEqual 判断两个 *int 指针是否相等
func IntPtrEqual(a, b *int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	return *a == *b
}
