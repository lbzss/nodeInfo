package node

import (
	"encoding/json"
	"github.com/shirou/gopsutil/v4/net"
)

type NetConfig struct {
	ProcessId   int32  `json:"process_id,omitempty"`
	LocalAddr   string `json:"local_addr,omitempty"`
	LocalPort   int32  `json:"local_port,omitempty"`
	RemoteAddr  string `json:"remote_addr,omitempty"`
	RemotePort  int32  `json:"remote_port,omitempty"`
	ProcessName string `json:"process_name,omitempty"`
	Status      string `json:"status,omitempty"`
}

func (n *NetConfig) String() string {
	data, _ := json.Marshal(n)
	return string(data)
}

var netTypes = []string{"tcp", "udp"}

func GetNetConfigData() ([]*NetConfig, map[int32][]*NetConfig, error) {
	var results []*NetConfig
	pidMap := make(map[int32][]*NetConfig)
	for _, netType := range netTypes {
		connections, err := net.Connections(netType)
		if err != nil {
			return nil, nil, err
		}
		for _, connection := range connections {
			result := &NetConfig{
				ProcessId:   connection.Pid,
				LocalAddr:   connection.Laddr.IP,
				LocalPort:   int32(connection.Laddr.Port),
				RemoteAddr:  connection.Raddr.IP,
				RemotePort:  int32(connection.Raddr.Port),
				ProcessName: "",
				Status:      connection.Status,
			}
			results = append(results, result)
			if ins, exist := pidMap[connection.Pid]; exist {
				pidMap[connection.Pid] = append(ins, result)
			} else {
				pidMap[connection.Pid] = []*NetConfig{result}
			}
		}
	}
	return results, pidMap, nil
}
