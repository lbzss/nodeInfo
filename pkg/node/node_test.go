package node

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	t.Run("node", func(t *testing.T) {
		n := NewNode()
		n.Collect()
		n.Complete()
		t.Log(n.String())
	})
}
