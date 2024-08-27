package node

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/process"
	"time"
)

type PsProcessConfig struct {
	Pid       int32
	Name      string
	StartTime string
	Envs      []string
	CmdLine   string
	Status    []string
	OpenFiles []string
}

func (p *PsProcessConfig) String() string {
	return fmt.Sprintf("%#v", p)
}

func GetProcessConfigData() ([]*PsProcessConfig, error) {
	var results []*PsProcessConfig
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}
	for _, processInfo := range processes {
		results = append(results, &PsProcessConfig{
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
		})
	}
	return results, err
}
