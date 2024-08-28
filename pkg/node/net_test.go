package node

import (
	"testing"
)

func TestGetNetConfigData(t *testing.T) {
	t.Run("net", func(t *testing.T) {
		netConfigs, netConfigsMap, _ := GetNetConfigData()
		for _, netConfig := range netConfigs {
			t.Log(netConfig.String())
		}
		t.Log(netConfigsMap)
	})
}
