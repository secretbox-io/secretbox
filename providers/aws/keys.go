package aws

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pkg/errors"
)

func (p AWSProvider) CreateMasterKey(path string) (string, error) {
	client := kms.New(p.cfg)

	req := client.CreateKeyRequest(&kms.CreateKeyInput{
		Description: aws.String("secretbox.io master key"),
		Tags: []kms.Tag{
			kms.Tag{
				TagKey:   aws.String("service"),
				TagValue: aws.String("secretbox.io"),
			},
			kms.Tag{
				TagKey:   aws.String("path"),
				TagValue: aws.String(path),
			},
		},
	})

	resp, err := req.Send()
	if err != nil {
		return "", errors.Wrap(err, "request to create master key failed")
	}
	if resp.KeyMetadata.KeyId == nil {
		return "", errors.New("failed to create new master key")
	}

	// create alias of path name
	alias := aliasPath(path)
	reqA := client.CreateAliasRequest(&kms.CreateAliasInput{
		AliasName:   aws.String(alias),
		TargetKeyId: resp.KeyMetadata.KeyId,
	})
	if _, err := reqA.Send(); err != nil {
		return "", errors.Wrap(err, "failed to create alias for master key")
	}

	return alias, nil
}

func (p AWSProvider) CheckKeyConfig(path string) (bool, error) {
	client := kms.New(p.cfg)
	req := client.DescribeKeyRequest(&kms.DescribeKeyInput{
		KeyId: aws.String(aliasPath(path)),
	})
	if _, err := req.Send(); err != nil {
		// TODO: check error types to distinguish different types of errors
		return false, nil
	}
	return true, nil
}

func aliasPath(path string) string {
	return strings.Join([]string{"alias", strings.Replace(path, "*", "_all_", -1)}, "/")
}

func unaliasPath(alias string) string {
	p := strings.SplitN(alias, "/", 2)
	return strings.Replace(p[1], "_all_", "*", -1)
}
