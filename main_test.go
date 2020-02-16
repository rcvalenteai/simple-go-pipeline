package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		a      int
		b      int
		answer int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{4, 0, 4},
		{-1, 3, 2},
	}

	for _, table := range tables {
		total := add(table.a, table.b)
		if total != table.answer {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d want: %d.", table.a, table.b, total, table.answer)
		}
	}
}
