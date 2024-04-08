package pointer

func ToPtr[T any](x T) *T {
	return &x
}
