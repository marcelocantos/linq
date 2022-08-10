package linq

import "math/rand"

// Sample returns a query that randomly samples each element in q with
// probability p. The returned query will deterministically sample values at the
// same intervals each time an enumerator is requested. This is not the case
// across calls to Sample.
func Sample[T any](q Query[T], p float64) Query[T] {
	return SampleSeed(q, p, rand.Int63())
}

// SampleSeed returns a query that randomly samples each element in q with
// probability p.
//
// The seed allows for deterministic results. Multiple invokations of
// SampleSeed with the same seed will return a query that samples values
// at the same intervals.
func SampleSeed[T any](q Query[T], p float64, seed int64) Query[T] {
	return NewQuery(func() Enumerator[T] {
		next := q.Enumerator()
		src := rand.NewSource(seed)
		rnd := rand.New(src)
		return func() (t T, ok bool) {
			for t, ok = next(); ok; t, ok = next() {
				if rnd.Float64() < p {
					return t, ok
				}
			}
			return t, ok
		}
	})
}
