package aws

import (
	"fmt"
	"log"

	"github.com/BTBurke/clt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/pkg/errors"
	"github.com/secretbox-io/secretbox/config"
)

type AWSProvider struct {
	cfg aws.Config
}

type creds struct {
	profile string
	access  string
	secret  string
	prompt  bool
}

type CredOption func(c *creds)

type ConfigProvider func() (*AWSProvider, error)

func NewProvider(opts ...CredOption) (*AWSProvider, error) {
	c := new(creds)
	for _, opt := range opts {
		opt(c)
	}

	// first check configured credentials and prompt if necessary
	if c.access != "" && c.secret != "" {
		p, err := configuredCredentials(access, secret)
		if err != nil {
			switch c.prompt {
			case true:
				access, secret := getCredsPrompt()
				return configuredCredentials(access, secret)
			default:
				return nil, errors.Wrap(err, "bad credentials in config file")
			}
		}
	}

	def, err := getDefaultCredentials()
	if err != nil {
		switch c.prompt {
		case true:

		}
	}

}

func getDefaultCredentials() ConfigProvider {
	return func() (*AWSProvider, error) {
		cfg, err := external.LoadDefaultAWSConfig()
		if err != nil {
			return nil, errors.Wrap(err, "could not load default config for AWS provider")
		}
		return &AWSProvider{
			cfg: cfg,
		}, nil
	}
}

func configuredCredentials(access string, secret string) (*AWSProvider, error) {
	return nil, nil
}

func WithProfile(profile string) CredOption {
	return func(c *creds) {
		c.profile = profile
	}
}

func WithCreds(access string, secret string) CredOption {
	return func(c *creds) {
		c.access = access
		c.secret = secret
	}
}

func Prompt(c *config.Config) CredOption {
	return func(c *creds) {
		i := clt.NewInteractiveSession()
		fmt.Println("Enter your AWS credentials")
		access := i.Ask("AWS Access Key ID: ")
		if access == "" {
			log.Fatal("Credentials are required to use secretbox")
		}
		secret := i.Ask("AWS Secret Key: ")
		if secret == "" {
			log.Fatal("Credentials are required to use secretbox")
		}
		c.access = access
		c.secret = secret

		c.Credentials = append(c.Credentials, config.Credentials{
			Profile: "default",
			Access:  access,
			Secret:  secret,
		})
		if err := c.Write(); err != nil {
			log.Fatal("failed to write credentials to config file")
		}
	}
}
