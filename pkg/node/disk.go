package node

import (
	"encoding/json"
	"github.com/shirou/gopsutil/v4/disk"
)

type DiskInfo struct {
	Device     string `json:"device,omitempty"`
	MountPoint string `json:"mount_point,omitempty"`
	FsType     string `json:"fs_type,omitempty"`
	Total      int64  `json:"total,omitempty"`
}

func (d *DiskInfo) String() string {
	data, _ := json.Marshal(d)
	return string(data)
}

func GetDiskData() ([]*DiskInfo, error) {
	var results []*DiskInfo
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	for _, partition := range partitions {
		results = append(results, &DiskInfo{
			Device:     partition.Device,
			MountPoint: partition.Mountpoint,
			FsType:     partition.Fstype,
			Total: func() int64 {
				stat, _ := disk.Usage(partition.Mountpoint)
				return int64(stat.Total)
			}(),
		})
	}
	return results, nil
}
