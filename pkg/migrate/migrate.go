package migrate

import (
	"fmt"
	"github.com/ek/pkg/aws"
	"github.com/ek/pkg/docker"
	"github.com/ek/pkg/domain"
	"regexp"
	"strings"
)

type Option struct {
	Namespace string
	Repository string
	Exclude string
	Config string
	Prefix string
}

type PullImageInfo struct {
	Status string `json:"status"`
	ProgressDetail map[string]string `json:"progressDetail,omitempty"`
	Progress string `json:"progress,omitempty"`
	Id string `json:"id"`
}

var regex *regexp.Regexp

func Run(images []string, config *Option) {

	if config.Exclude != "" {
		regex = regexp.MustCompile(config.Exclude)
		if regex == nil {
			fmt.Errorf("regex error")
			return
		}
	}

	ecr := aws.NewEcr()
	dockerCli := docker.NewClient(config.Config)

	for _, image := range images {
		if config.Exclude != "" {
			result := regex.FindAllStringSubmatch(image, -1)
			if result != nil {
				continue
			}
		}

		// 迁移镜像
		// abc.io/mammoth/promtail:2.2.0
		parts := strings.Split(image, ":")
		// version = 2.2.0
		version := parts[1]
		// abc.io/mammoth/promtail
		repositoryParts := strings.Split(parts[0], "/")
		var tagName = image
		var repository string
		if domain.IsDomain(repositoryParts[0]) {
			tagName = strings.TrimPrefix(image, repositoryParts[0] + "/")
		}
		repository = strings.TrimSuffix(tagName, ":" + version)
		if config.Prefix != "" {
			repository = config.Prefix +  "/" + repository
			tagName = config.Prefix + "/" + tagName
		}
		tagName = config.Repository + "/" + tagName

		// 检查镜像是否存在
		imageExist := ecr.CheckImageIsExist(repository, version)
		if !imageExist {
			dockerCli.PullImage(image)
			fmt.Printf("tag image : %v -> %v\n", image, tagName)
			dockerCli.ImageTag(image, tagName)
			dockerCli.PushImage(tagName)
		} else {
			fmt.Printf("镜像 [%v] 已经存在\n", image)
		}
	}
}