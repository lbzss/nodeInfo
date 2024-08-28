package node

import (
	"encoding/json"
	"github.com/shirou/gopsutil/v4/net"
)

type InterfaceConfig struct {
	Name string   `json:"name,omitempty"`
	Addr []string `json:"addr,omitempty"`
}

func (i *InterfaceConfig) String() string {
	data, _ := json.Marshal(i)
	return string(data)
}

func GetInterfaceData() ([]*InterfaceConfig, error) {
	var results []*InterfaceConfig

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, interfaceInfo := range interfaces {
		var addrInfo []string
		addrs := interfaceInfo.Addrs
		for _, addr := range addrs {
			addrInfo = append(addrInfo, addr.Addr)
		}
		results = append(results, &InterfaceConfig{
			Name: interfaceInfo.Name,
			Addr: addrInfo,
		})
	}

	return results, nil
}
