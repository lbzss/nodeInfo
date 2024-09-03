package node

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"strings"
)

var DockerClient *client.Client

type DockerContainer struct {
	ContainerId string `json:"container_id,omitempty"`
	Pid         int64  `json:"pid,omitempty"`
	// todo: returns the md5 value currently. the image should get tht repository/image_name,
	Image   string       `json:"image,omitempty"`
	Command string       `json:"command,omitempty"`
	Mounts  []MountPoint `json:"mounts,omitempty"`
	// todo: also get the port map
}

type Port struct {
	IP          string `json:"ip,omitempty"`
	PrivatePort int    `json:"private_port,omitempty"`
	PublicPort  int    `json:"public_port,omitempty"`
	Type        string `json:"type,omitempty"`
}

type MountPoint struct {
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Source      string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
	Mode        string `json:"mode,omitempty"`
	RW          bool   `json:"rw,omitempty"`
}

// GetContainerImages get a map from containerImage id to containerImage RepoTag
func GetContainerImages() (map[string]string, error) {
	result := make(map[string]string)
	containerImages, err := DockerClient.ImageList(context.Background(), image.ListOptions{
		All: true,
	})

	if err != nil {
		return nil, err
	}
	for _, containerImage := range containerImages {
		result[containerImage.ID] = strings.Join(containerImage.RepoTags, " ")
	}
	return result, nil
}

func GetContainerData() ([]*DockerContainer, map[int32]*DockerContainer, error) {
	var results []*DockerContainer
	dockerContainerMap := make(map[int32]*DockerContainer)
	ctx := context.Background()

	imageMap, err := GetContainerImages()
	if err != nil {
		return nil, nil, err
	}
	containers, err := DockerClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return nil, nil, err
	}
	for _, dockerContainer := range containers {
		//todo: should we skip if there's error.
		data, err := DockerClient.ContainerInspect(ctx, dockerContainer.ID)
		if err != nil {
			return nil, nil, err
		}
		result := &DockerContainer{
			ContainerId: data.ID,
			Pid:         int64(data.State.Pid),
			Image:       imageMap[data.Image],
			Command:     strings.Join(data.Args, ","),
			//Ports:       nil,
			Mounts: func() []MountPoint {
				var res []MountPoint
				for _, v := range data.Mounts {
					res = append(res, MountPoint{
						Type:        string(v.Type),
						Name:        v.Name,
						Source:      v.Source,
						Destination: v.Destination,
						Mode:        v.Mode,
						RW:          v.RW,
					})
				}
				return res
			}(),
		}
		results = append(results, result)
		if _, exist := dockerContainerMap[int32(result.Pid)]; !exist {
			dockerContainerMap[int32(result.Pid)] = result
		}
	}
	return results, dockerContainerMap, nil
}

func CreateClient() error {
	if DockerClient != nil {
		return nil
	}
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.43"))
	if err != nil {
		return err
	}
	DockerClient = dockerClient
	return nil
}

func init() {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.43"))
	if err != nil {
		panic(err)
	}
	DockerClient = dockerClient
}
