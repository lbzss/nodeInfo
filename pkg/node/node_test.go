package node

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	t.Run("node", func(t *testing.T) {
		n := NewNode()
		n.Collect()
		t.Logf("%+v", n)
	})
}
