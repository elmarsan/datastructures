package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("push", func(t *testing.T) {
		sp := NewStack()

		var val uint16 = 25
		sp.push(val)

		if *sp.last != val {
			t.Error("Unable to push val")
		}

		if sp.stack[0] != val {
			t.Error("Unable to push val")
		}
	})

	t.Run("pop with one element", func(t *testing.T) {
		sp := NewStack()

		var val uint16 = 25
		sp.push(val)

		last := sp.pop()

		if *last != val {
			t.Error("Unable to pop val")
		}

		if len(sp.stack) != 0 {
			t.Error("Unable to pop val")
		}

		if sp.last != nil {
			t.Error("Unable to pop val")
		}
	})

	t.Run("pop with more than one element", func(t *testing.T) {
		sp := NewStack()

		var val uint16 = 25
		sp.push(val)
		sp.push(val)

		last := sp.pop()

		if *last != val {
			t.Error("Unable to pop val")
		}

		if len(sp.stack) != 1 {
			t.Error("Unable to pop val")
		}

		if sp.last == nil {
			t.Error("Unable to pop val")
		}
	})

	t.Run("pop empty stack", func(t *testing.T) {
		sp := NewStack()

		last := sp.pop()

		if last != nil {
			t.Error("Wrong pop")
		}
	})
}
