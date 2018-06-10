package aws

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pkg/errors"
	"github.com/secretbox-io/secretbox/crypto"
)

// CreateMasterKey creates a customer master key in AWS at the given alias path.  Paths are of the form
// secretbox/<project>/<environment>.  The default key is located at path secretbox/*/production.
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

// KeyExists checks to see if the key located at path exists
func (p AWSProvider) KeyExists(path string) (bool, error) {
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

// GenerateDataKeyContext uses the customer master key associated with keyID to generate a unique data
// encryption key.  It returns the unencrypted data key and the encrypted key as a keyset.  This method can
// be cancelled with a custom context value.
func (p AWSProvider) GenerateDataKeyContext(ctx context.Context, keyID string) (*crypto.Keyset, error) {
	client := kms.New(p.cfg)
	req := client.GenerateDataKeyRequest(&kms.GenerateDataKeyInput{
		KeyId:   aws.String(aliasPath(keyID)),
		KeySpec: kms.DataKeySpecAes256,
	})
	req.SetContext(ctx)
	resp, err := req.Send()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create DEK")
	}
	return &crypto.Keyset{
		DEK:          resp.Plaintext,
		EncryptedDEK: resp.CiphertextBlob,
	}, nil
}

// GenerateDataKey uses the customer master key associated with keyID to generate a unique data
// encryption key.  It returns the unencrypted data key and the encrypted key as a keyset.
func (p AWSProvider) GenerateDataKey(keyID string) (*crypto.Keyset, error) {
	return p.GenerateDataKeyContext(context.Background(), keyID)
}

func aliasPath(path string) string {
	return strings.Join([]string{"alias", strings.Replace(path, "*", "_all_", -1)}, "/")
}

func unaliasPath(alias string) string {
	p := strings.SplitN(alias, "/", 2)
	return strings.Replace(p[1], "_all_", "*", -1)
}
