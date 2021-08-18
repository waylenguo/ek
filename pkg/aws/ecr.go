package aws

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"strings"
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

func (cli *Ecr) CreateRepository(repository string) {
	context := context.TODO()

	svc := cli.instance

	_, err := svc.CreateRepository(context, &ecr.CreateRepositoryInput{
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

func (cli *Ecr) GetLoginPassword() *string {
	svc := cli.instance
	output, err := svc.GetAuthorizationToken(context.TODO(), &ecr.GetAuthorizationTokenInput{})
	if err != nil {
		panic(err)
	}
	auth := output.AuthorizationData[0]
	authTokenByte, err := base64.StdEncoding.DecodeString(*auth.AuthorizationToken)
	authToken := string(authTokenByte)
	authToken = strings.Split(authToken, ":")[1]
	return &authToken
}