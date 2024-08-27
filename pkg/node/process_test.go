package node

import (
	"testing"
)

func TestGetProcessData(t *testing.T) {

	t.Run("process", func(t *testing.T) {
		processes, _ := GetProcessConfigData()
		for _, process := range processes {
			t.Log(process.String())
		}
	})
}
