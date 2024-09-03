package node

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	HostConfig         *HostInfo                  `json:"host_config,omitempty"`
	NetConfigs         []*NetConfig               `json:"net_configs,omitempty"`
	NetConfigsMap      map[int32][]*NetConfig     `json:"-"`
	Processes          []*PsProcessConfig         `json:"process_configs,omitempty"`
	ProcessesMap       map[int32]*PsProcessConfig `json:"-"`
	Disk               []*DiskInfo                `json:"disk_configs,omitempty"`
	Interfaces         []*InterfaceConfig         `json:"interfaces,omitempty"`
	DockerContainers   []*DockerContainer         `json:"docker_containers,omitempty"`
	DockerContainerMap map[int32]*DockerContainer `json:"-"`
}

func NewNode() *Node {
	return &Node{
		HostConfig:       &HostInfo{},
		NetConfigs:       make([]*NetConfig, 0),
		Processes:        make([]*PsProcessConfig, 0),
		Disk:             make([]*DiskInfo, 0),
		Interfaces:       make([]*InterfaceConfig, 0),
		DockerContainers: make([]*DockerContainer, 0),
	}
}

func (n *Node) String() string {
	data, _ := json.Marshal(n)
	return string(data)
}

func (n *Node) Collect() []error {
	var errorList []error
	hostConfig, err := GetHostData()
	if err != nil {
		errorList = append(errorList, err)
	}
	n.HostConfig = hostConfig

	netConfigs, netConfigsMap, err := GetNetConfigData()
	if err != nil {
		errorList = append(errorList, err)
	}
	n.NetConfigs = netConfigs
	n.NetConfigsMap = netConfigsMap

	psProcessConfigs, psProcessConfigMap, err := GetProcessConfigData()
	if err != nil {
		errorList = append(errorList, err)
	}
	n.Processes = psProcessConfigs
	n.ProcessesMap = psProcessConfigMap

	diskInfos, err := GetDiskData()
	if err != nil {
		errorList = append(errorList, err)
	}
	n.Disk = diskInfos

	interfaceData, err := GetInterfaceData()
	if err != nil {
		errorList = append(errorList, err)
	}
	n.Interfaces = interfaceData

	dockerContainerInfo, dockerContainerMap, err := GetContainerData()
	if err != nil {
		errorList = append(errorList, err)
	}
	n.DockerContainers = dockerContainerInfo
	n.DockerContainerMap = dockerContainerMap
	return errorList
}

func (n *Node) Complete() {
	if n.ProcessesMap == nil {
		fmt.Println("ProcessesMap empty")
		return
	}
	for _, netConfig := range n.NetConfigs {
		if process, exists := n.ProcessesMap[netConfig.ProcessId]; exists && process != nil {
			netConfig.ProcessName = process.Name
		} else {
			netConfig.ProcessName = "unknown"
		}
	}
}
