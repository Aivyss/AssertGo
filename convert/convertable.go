package convert

func IsConvertable[T any](obj interface{}) bool {
	_, ok := obj.(T)

	return ok
}
