package linq

// Identity returns t unmodified.
func Identity[T any](t T) T {
	return t
}

func drain[T any](next Enumerator[T]) {
	for _, ok := next(); ok; {
		_, ok = next()
	}
}

func valueElse[T any](t T, ok bool, alt T) T { //nolint:revive
	if ok {
		return t
	}
	return alt
}

func counter(i int) func() int {
	i--
	return func() int {
		i++
		return i
	}
}

func indexify[T, U any](f func(t T) U) func(i int, t T) U {
	return func(i int, t T) U {
		return f(t)
	}
}
