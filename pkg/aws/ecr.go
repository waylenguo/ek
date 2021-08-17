package aws

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

type Ecr struct {
	instance *ecr.Client
}

func NewEcr() *Ecr {

	context := context.TODO()
	cfg, err := config.LoadDefaultConfig(context)
	if err != nil {
		fmt.Errorf("unable to load SDK config, %v", err)
	}

	return &Ecr{
		instance: ecr.NewFromConfig(cfg),
	}
}

func (*Ecr) CreateRepository(repository string) {
	context := context.TODO()
	cfg, err := config.LoadDefaultConfig(context)
	if err != nil {
		fmt.Errorf("unable to load SDK config, %v", err)
	}

	svc := ecr.NewFromConfig(cfg)
	_, err = svc.CreateRepository(context, &ecr.CreateRepositoryInput{
		RepositoryName: &repository,
	})
	if err != nil {
		var rae *types.RepositoryAlreadyExistsException
		if errors.As(err, &rae) {
			fmt.Printf("The repository with name '%s' already exist\n", repository)
		} else {
			panic(err)
		}
	}

}