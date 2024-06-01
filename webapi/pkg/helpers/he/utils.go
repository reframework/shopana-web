package he

func WithDefault[T comparable](value T, defaultValue T) T {
	var v T

	if value == v {
		return defaultValue
	}

	return value
}
