package node

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/net"
)

type NetConfig struct {
	ProcessId   int32
	Port        int32
	ProcessName string
	Status      string
}

func (n *NetConfig) String() string {
	return fmt.Sprintf("%#v", n)
}

var netTypes = []string{"tcp", "udp"}

func GetNetConfigData() ([]*NetConfig, error) {
	var results []*NetConfig
	for _, netType := range netTypes {
		connections, err := net.Connections(netType)
		if err != nil {
			return nil, err
		}
		for _, connection := range connections {
			results = append(results, &NetConfig{
				ProcessId:   connection.Pid,
				Port:        int32(connection.Laddr.Port),
				ProcessName: "",
				Status:      connection.Status,
			})
		}
	}
	return results, nil
}
