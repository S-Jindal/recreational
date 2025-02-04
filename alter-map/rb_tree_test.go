package rbtree

import (
	"math/rand"
	"testing"
)

func TestRbTree(t *testing.T) {
	tree := &RbTree{}

	// Test empty tree
	if tree.Get(1) != nil {
		t.Error("Expected nil for non-existent key")
	}

	// Test single insert
	tree.Put(1, "one")
	if tree.Get(1) != "one" {
		t.Error("Expected 'one' for key 1")
	}
	if tree.Length != 1 {
		t.Error("Expected length 1")
	}
	if tree.root.nodeColor != black {
		t.Error("Root should be black")
	}

	// Test multiple inserts
	tree.Put(2, "two")
	tree.Put(3, "three")

	if tree.Get(2) != "two" {
		t.Error("Expected 'two' for key 2")
	}
	if tree.Get(3) != "three" {
		t.Error("Expected 'three' for key 3")
	}
	if tree.Length != 3 {
		t.Error("Expected length 3")
	}

	// Test update existing key
	tree.Put(2, "TWO")
	if tree.Get(2) != "TWO" {
		t.Error("Expected 'TWO' for key 2 after update")
	}
	if tree.Length != 3 {
		t.Error("Length should not change after update")
	}

	tree.Put(4, "four")
	tree.Put(5, "five")
	tree.Put(6, "six")

	if tree.Get(4) != "four" {
		t.Error("Expected 'four' for key 4")
	}
	if tree.Get(5) != "five" {
		t.Error("Expected 'five' for key 5")
	}
	if tree.Get(6) != "six" {
		t.Error("Expected 'six' for key 6")
	}
	if tree.Length != 6 {
		t.Error("Expected length 6")
	}
	// Test illegal Get
	if tree.Get(7) != nil {
		t.Error("Expected nil for non-existent key 7")
	}
	if tree.Get(-1) != nil {
		t.Error("Expected nil for non-existent key -1")
	}
	// Test properties
	validateRedBlackProperties(t, tree.root)

}

// Helper function to validate Red-Black tree properties
func validateRedBlackProperties(t *testing.T, n *node) {
	if n == nil {
		return
	}

	// Property 1: Every node is either red or black
	if n.nodeColor != red && n.nodeColor != black {
		t.Error("Invalid color found")
	}

	// Property 2: Root is black (checked in main test)

	// Property 3: Red nodes should have black children
	if n.nodeColor == red {
		if n.left != nil && n.left.nodeColor == red {
			t.Error("Red node has red left child")
		}
		if n.right != nil && n.right.nodeColor == red {
			t.Error("Red node has red right child")
		}
	}

	// Recursively check children
	validateRedBlackProperties(t, n.left)
	validateRedBlackProperties(t, n.right)
}

func BenchmarkRbTree(b *testing.B) {
	rng := rand.New(rand.NewSource(4))
	datasetSize := 100

	// Pre-generate random keys for consistent benchmark conditions
	insertKeys := make([]int, datasetSize)
	lookupKeys := make([]int, datasetSize)

	for i := 0; i < datasetSize; i++ {
		insertKeys[i] = rng.Int()
		lookupKeys[i] = rng.Int()
	}

	b.Log("Running tests on dataset size:", datasetSize)

	b.Run("RbTree-Sequential-Insert", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tree := &RbTree{}
			for i := 0; i < datasetSize; i++ {
				tree.Put(insertKeys[i], struct{}{})
			}
		}
	})

	b.Run("RbTree-Sequential-Lookup", func(b *testing.B) {
		// First build the tree
		tree := &RbTree{}
		for i := 0; i < datasetSize; i++ {
			tree.Put(insertKeys[i], struct{}{})
		}

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			for i := 0; i < datasetSize; i++ {
				_ = tree.Get(lookupKeys[i])
			}
		}
	})

	b.Run("RbTree-Mixed-Operations", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			tree := &RbTree{}
			for i := 0; i < datasetSize; i++ {
				tree.Put(insertKeys[i], struct{}{})
				_ = tree.Get(lookupKeys[i])
			}
		}
	})

	b.Run("Map-Sequential-Insert", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			m := make(map[int]struct{})
			for i := 0; i < datasetSize; i++ {
				m[insertKeys[i]] = struct{}{}
			}
		}
	})

	b.Run("Map-Sequential-Lookup", func(b *testing.B) {
		// First build the map
		m := make(map[int]struct{})
		for i := 0; i < datasetSize; i++ {
			m[insertKeys[i]] = struct{}{}
		}

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			for i := 0; i < datasetSize; i++ {
				_ = m[lookupKeys[i]] // map lookup is direct and fast
			}
		}
	})

	b.Run("Map-Mixed-Operations", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			m := make(map[int]struct{})
			for i := 0; i < datasetSize; i++ {
				m[insertKeys[i]] = struct{}{}
				_ = m[lookupKeys[i]] // map lookup is direct and fast
			}
		}
	})
}
