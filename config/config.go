package config

import (
	"os"
	"path/filepath"

	"github.com/BTBurke/go-homedir"
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Provider is an enum type for supported cloud platforms
type Provider string

const (
	// AWS is a supported cloud provider
	AWS Provider = "aws"
)

// T represents the TOML configuration, not all of which may be set.  Prefer Config with contains an embedded ConfigT
// with additional metadata about which values were set in the file.
type T struct {
	Email    string
	Token    string
	Provider Provider
	Keys     []Key
}

// Key is a customer master key in the could providers KMS.  ID should be sufficient to look up this particular key, with optional
// path to indicate what environment this key would be used in, of the form `secretbox/<project>/<environment>`.  For example, the first
// master key created is `secretbox/*/production` and is used for all projects unless a different key is configure on a per-project basis.
type Key struct {
	ID   string
	Path string
}

// Config represents the configuration with additional metadata about which values were set in the file.
type Config struct {
	T
	MetaData toml.MetaData
}

// Writer is a configuration that may be updated and then written to the file
type Writer interface {
	Write() error
}

// New writes a minimal configuration to the file
func New(email string, token string, provider Provider, keys ...Key) (*Config, error) {
	cfg := new(Config)
	cfg.Email = email
	cfg.Token = token
	cfg.Provider = provider
	cfg.Keys = keys

	if err := cfg.Write(); err != nil {
		return nil, errors.Wrap(err, "could not write initial config")
	}
	// read back config to populate metadata
	return Read()
}

// Read will read the config file located at `%HOME%/.secretbox/config.toml` and return a populated Config
// and metadata bout what values were explicitly set.`
func Read() (*Config, error) {
	cfgFile, err := getConfigFileName()
	if err != nil {
		log.Fatal("could not find home directory")
	}

	var cfg T
	md, err := toml.DecodeFile(cfgFile, &cfg)
	if err != nil {
		log.Fatal("could not read configuration file.  Try running `secretbox login` first.")
	}

	return &Config{
		cfg,
		md,
	}, nil
}

func (c *Config) Write() error {
	cfgFile, err := getConfigFileName()
	if err != nil {
		return errors.Wrap(err, "could not get config file name")
	}
	cfgPath, err := getConfigFilePath()
	if err != nil {
		return errors.Wrap(err, "could not get config file path")
	}
	if err := os.MkdirAll(cfgPath, os.ModeDir); err != nil {
		return errors.Wrap(err, "could not create $HOME/.secretbox directory")
	}

	f, err := os.OpenFile(cfgFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return errors.Wrap(err, "could not open config file for writing")
	}
	if err := toml.NewEncoder(f).Encode(c.T); err != nil {
		return errors.Wrap(err, "could not write to config file")
	}
	if err := f.Close(); err != nil {
		return errors.Wrap(err, "could not close config file")
	}
	return nil
}

func getConfigFileName() (string, error) {
	p, err := getConfigFilePath()
	if err != nil {
		return "", err
	}
	cfgFile := filepath.Join(p, "config.toml")
	return cfgFile, nil
}

func getConfigFilePath() (string, error) {
	home, err := getHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".secretbox"), nil
}

func getHome() (string, error) {
	homedir.WinPreferUserProfile = true
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	expand, err := homedir.Expand(home)
	if err != nil {
		return "", err
	}
	return expand, nil
}

// GetHomeDir returns the home directory based on OS type
func GetHomeDir() (string, error) {
	return getHome()
}
