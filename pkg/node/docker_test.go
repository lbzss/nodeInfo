package node

import (
	"testing"
)

func TestGetContainerData(t *testing.T) {
	t.Run("dockerContainer", func(t *testing.T) {
		infos, dockerContainerMap, _ := GetContainerData()
		for _, info := range infos {
			t.Logf("%+v", info)
		}
		t.Log(dockerContainerMap)
	})
}

func TestGetContainerImages(t *testing.T) {
	t.Run("containerImages", func(t *testing.T) {
		infos, _ := GetContainerImages()
		for _, info := range infos {
			t.Logf("%+v", info)
		}
		//t.Log(dockerContainerMap)
	})
}
