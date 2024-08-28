package node

import (
	"testing"
)

func TestGetInterfaceData(t *testing.T) {
	t.Run("interface", func(t *testing.T) {
		interfaceConfigs, _ := GetInterfaceData()
		for _, netConfig := range interfaceConfigs {
			t.Log(netConfig)
		}
	})
}
