package crypto

import "context"

// Keyset represents a pair of a data encryption key and the associated DEK encrypted
// with a customer master key
type Keyset struct {
	DEK          []byte
	EncryptedDEK []byte
}

// EncryptionKeyProvider is fulfilled by a provider to generate data encryption keys given a customer master key
// specified by keyID
type EncryptionKeyProvider interface {
	GenerateDataKey(keyID string) (*Keyset, error)
	GenerateDataKeyContext(ctx context.Context, keyID string) (*Keyset, error)
}

// DecryptionKeyProvider is fulfilled by a provider to decrypt data encryption keys given a customer master key
// specified by keyID
type DecryptionKeyProvider interface {
	DecryptDataKey(encryptedKey []byte, keyID string) ([]byte, error)
	DecryptDataKeyContext(ctx context.Context, encryptedKey []byte, keyID string) ([]byte, error)
}

// KeyManager is fulfilled by a provider to create and manage customer master keys in the cloud key management service
type KeyManager interface {
	KeyExists() (bool, error)
	CreateMasterKey(path string) (string, error)
	//DeleteMasterKey(path string) error
}

// RoleManager provides create, delete, and verify role-based access to KMS key material for cloud infrastructure
type RoleManager interface {
	CheckRoleConfig() (bool, error)
	CreateRole() (string, error)
	DeleteRole() error
}
