package node

import (
	"encoding/json"
	"github.com/shirou/gopsutil/v4/process"
	"time"
)

type PsProcessConfig struct {
	Pid       int32    `json:"pid,omitempty"`
	Name      string   `json:"name,omitempty"`
	StartTime string   `json:"start_time,omitempty"`
	Envs      []string `json:"envs,omitempty"`
	CmdLine   string   `json:"cmd_line,omitempty"`
	Status    []string `json:"status,omitempty"`
	OpenFiles []string `json:"open_files,omitempty"`
}

func (p *PsProcessConfig) String() string {
	data, _ := json.Marshal(p)
	return string(data)
}

func GetProcessConfigData() ([]*PsProcessConfig, map[int32]*PsProcessConfig, error) {
	var results []*PsProcessConfig
	pidMap := make(map[int32]*PsProcessConfig)
	processes, err := process.Processes()
	if err != nil {
		return nil, nil, err
	}
	for _, processInfo := range processes {
		result := &PsProcessConfig{
			Pid: processInfo.Pid,
			Name: func() string {
				name, err := processInfo.Name()
				if err != nil {
					return "UNKNOWN"
				}
				return name
			}(),
			StartTime: func() string {
				createTime, err := processInfo.CreateTime()
				if err != nil {
					return "UNKNOWN"
				}
				return time.Unix(createTime, 0).Format("2006-01-02 15:04:05")
			}(),
			Envs: func() []string {
				envs, err := processInfo.Environ()
				if err != nil {
					return []string{}
				}
				return envs
			}(),
			CmdLine: func() string {
				cmdLine, err := processInfo.Cmdline()
				if err != nil {
					return "UNKNOWN"
				}
				return cmdLine
			}(),
			Status: func() []string {
				status, err := processInfo.Status()
				if err != nil {
					return []string{}
				}
				return status
			}(),
			OpenFiles: func() []string {
				var results []string
				openFiles, err := processInfo.OpenFiles()
				if err != nil {
					return []string{}
				}
				for _, openFile := range openFiles {
					results = append(results, openFile.Path)
				}
				return results
			}(),
		}
		results = append(results, result)
		if _, exist := pidMap[result.Pid]; !exist {
			pidMap[result.Pid] = result
		}

	}
	return results, pidMap, err
}
