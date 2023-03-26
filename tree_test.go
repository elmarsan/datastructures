package main

import (
	"testing"
)

func TestTree(t *testing.T) {
	tree := &Tree{
		v: 1,
		r: &Tree{
			v: 3,
		},
		l: &Tree{
			v: 2,
			l: &Tree{
				v: 7,
			},
			r: &Tree{
				v: 1,
			},
		},
	}

	t.Run("depth", func(t *testing.T) {
		if tree.depth() != 3 {
			t.Error("Wrong depth")
		}
	})

	t.Run("sum", func(t *testing.T) {
		if tree.Sum() != 14 {
			t.Error("Wrong sum")
		}
	})
}
