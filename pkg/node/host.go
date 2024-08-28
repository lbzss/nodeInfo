package node

import (
	"encoding/json"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

type HostInfo struct {
	Hostname             string `json:"hostname,omitempty"`
	Uptime               int64  `json:"uptime,omitempty"`
	BootTime             int64  `json:"boot_time,omitempty"`
	Procs                int64  `json:"procs,omitempty"`
	Cores                int64  `json:"cores,omitempty"`
	MemTotal             int64  `json:"mem_total,omitempty"`
	OS                   string `json:"os,omitempty"`
	Platform             string `json:"platform,omitempty"`
	PlatformFamily       string `json:"platform_family,omitempty"`
	PlatformVersion      string `json:"platform_version,omitempty"`
	KernelVersion        string `json:"kernel_version,omitempty"`
	KernelArch           string `json:"kernel_arch,omitempty"`
	VirtualizationSystem string `json:"virtualization_system,omitempty"`
	VirtualizationRole   string `json:"virtualization_role,omitempty"`
	HostID               string `json:"host_id,omitempty"`
}

func (h *HostInfo) String() string {
	data, _ := json.Marshal(h)
	return string(data)
}

func GetHostData() (*HostInfo, error) {
	info, err := host.Info()
	if err != nil {
		return nil, err
	}
	return &HostInfo{
		Hostname: info.Hostname,
		Uptime:   int64(info.Uptime),
		BootTime: int64(info.BootTime),
		Procs:    int64(info.Procs),
		Cores: func() int64 {
			count, _ := cpu.Counts(true)
			return int64(count)
		}(),
		MemTotal: func() int64 {
			memInfo, _ := mem.VirtualMemory()
			return int64(memInfo.Total) / 1024 / 1024
		}(),
		OS:                   info.OS,
		Platform:             info.Platform,
		PlatformFamily:       info.PlatformFamily,
		PlatformVersion:      info.PlatformVersion,
		KernelVersion:        info.KernelVersion,
		KernelArch:           info.KernelArch,
		VirtualizationSystem: info.VirtualizationSystem,
		VirtualizationRole:   info.VirtualizationRole,
		HostID:               info.HostID,
	}, nil
}
