package node

import (
	"testing"
)

func TestGetDiskData(t *testing.T) {
	t.Run("disk", func(t *testing.T) {
		info, _ := GetDiskData()
		t.Log(info)
	})
}
