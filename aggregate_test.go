package linq_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

func add(a, b int) int { return a + b }

func weight(a float64, i int, b float64) float64 {
	return a + float64(i)*b
}

func TestAggregate(t *testing.T) {
	t.Parallel()

	assertNoResult(t, maybe(linq.Iota1(0).Aggregate(add)))
	assertResultEqual(t, 15, maybe(linq.Iota2(1, 6).Aggregate(add)))
}

func TestAggregateElse(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 15, linq.From(1, 2, 3, 4, 5).AggregateElse(add, 42))

	assert.Equal(t, 42, linq.From[int]().AggregateElse(add, 42))
}

func TestAggregateSeed(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 42+15,
		linq.From(1, 2, 3, 4, 5).
			AggregateSeed(42, func(a, b int) int { return a + b }),
	)

	assert.Equal(t, 42,
		linq.From[int]().
			AggregateSeed(42, func(a, b int) int { return a + b }),
	)

	assert.Equal(t, ".1.2.3.4.5",
		linq.AggregateSeed(linq.From(1, 2, 3, 4, 5), "",
			func(a string, b int) string { return fmt.Sprintf("%s.%d", a, b) },
		),
	)
}

func TestAggregateI(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1.25, linq.From(0.25, 0.25, 0.5).AggregateSeedI(0, weight))
}

func TestMustAggregate(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 15, linq.From(1, 2, 3, 4, 5).MustAggregate(add))

	assert.PanicsWithError(t, "empty source",
		func() { linq.From[int]().MustAggregate(add) },
	)
}
