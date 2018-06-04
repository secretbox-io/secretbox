package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/pkg/errors"
)

type AWSProvider struct {
	cfg aws.Config
}

func NewProvider() (*AWSProvider, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return nil, errors.Wrap(err, "could not load default config for AWS provider")
	}
	return &AWSProvider{
		cfg: cfg,
	}, nil
}
