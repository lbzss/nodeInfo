package node

import "testing"

func TestCreateClient(t *testing.T) {

	t.Run("docker", func(t *testing.T) {
		CreateClient()
	})
}
