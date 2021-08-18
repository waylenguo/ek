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
	Image string
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

func MigrateImage(images []string, config *Option) {

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

		fmt.Printf("pulling image: %v\n", image)


		dockerCli.PullImage(image)
		// 迁移镜像
		tag, res := getRepositoryName(image, config)
		tag = config.Repository + tag
		fmt.Printf("tag image : %v -> %v\n", image, tag)

		dockerCli.ImageTag(image, tag)

		// 推送镜像
		ecr.CreateRepository(res)
		fmt.Println("==============================================================================================")
		dockerCli.PushImage(tag)
	}
}


func getRepositoryName(image string, config *Option) (tagName string, repositoryName string) {
	parts := strings.Split(image, "/")
	var (
		tag string
		repository string
	)
	for i, part := range parts {
		if i == 0 && domain.IsDomain(part) {
			continue
		} else {
			tag = tag + "/" + part
			repository = repository + "/" + part
		}
	}
	repository = strings.Split(repository, ":")[0]

	if config.Prefix != "" {
		tag = "/" + config.Prefix + tag
		repository = config.Prefix + repository
	}

	return tag, repository
}

