package node

import (
	"testing"
)

func TestGetHostData(t *testing.T) {
	t.Run("host", func(t *testing.T) {
		info, _ := GetHostData()
		t.Log(info)
	})
}
