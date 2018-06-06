package providers

import "github.com/secretbox-io/secretbox/crypto"

// Provider is the top interface for any cloud provider that implements the full suite of interfaces to provide
// the secretbox functionality
type Provider interface {
	crypto.Encrypter
}
