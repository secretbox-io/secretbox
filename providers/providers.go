package providers

import "github.com/secretbox-io/secretbox/crypto"

// Provider is the top interface for any cloud provider that implements the full suite of interfaces to provide
// the secretbox functionality
type Provider interface {
	crypto.Encrypter
}

// KeyManager is fulfilled by a provider to create and manage customer master keys in the cloud key management service
type KeyManager interface {
	CheckKeyConfig() (bool, error)
	CreateMasterKey(path string) (string, error)
	DeleteMasterKey(path string) error
}

// RoleManager provides create, delete, and verify role-based access to KMS key material for cloud infrastructure
type RoleManager interface {
	CheckRoleConfig() (bool, error)
	CreateRole() (string, error)
	DeleteRole() error
}

// CredentialsManager provides access to user credentials
type CredentialsManager interface {
	GetCredentials() error
}
