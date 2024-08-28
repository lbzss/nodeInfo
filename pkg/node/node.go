package node

type Node struct {
	HostConfig    *HostInfo                    `json:"host_config,omitempty"`
	NetConfigs    []*NetConfig                 `json:"net_configs,omitempty"`
	NetConfigsMap map[int32][]*NetConfig       `json:"net_configs_map,omitempty"`
	Processes     []*PsProcessConfig           `json:"process_configs,omitempty"`
	ProcessesMap  map[int32][]*PsProcessConfig `json:"processes_map,omitempty"`
	Disk          []*DiskInfo                  `json:"disk_configs,omitempty"`
	Interfaces    []*InterfaceConfig           `json:"interfaces,omitempty"`
}

func NewNode() *Node {
	return &Node{
		HostConfig: &HostInfo{},
		NetConfigs: make([]*NetConfig, 0),
		Processes:  make([]*PsProcessConfig, 0),
		Disk:       make([]*DiskInfo, 0),
		Interfaces: make([]*InterfaceConfig, 0),
	}
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
	return errorList
}
