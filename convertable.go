package assert

func ErrorTypeIs[T any](obj interface{}) bool {
	_, ok := obj.(T)

	return ok
}
