package node

import (
	"testing"
)

func TestGetNetConfigData(t *testing.T) {
	t.Run("net", func(t *testing.T) {
		netConfigs, _ := GetNetConfigData()
		for _, netConfig := range netConfigs {
			t.Log(netConfig.String())
		}
	})
}
