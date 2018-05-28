package crypto

import "context"

// EncryptionKeyProvider is fulfilled by a provider to generate data encryption keys given a customer master key
// specified by keyID
type EncryptionKeyProvider interface {
	GenerateDataKey(keyID string) ([]byte, error)
	GenerateDataKeyContext(ctx context.Context, keyID string) ([]byte, error)
}

// DecryptionKeyProvider is fulfilled by a provider to decrypt data encryption keys given a customer master key
// specified by keyID
type DecryptionKeyProvider interface {
	DecryptDataKey(encryptedKey []byte, keyID string) ([]byte, error)
	DecryptDataKeyContext(ctx context.Context, encryptedKey []byte, keyID string) ([]byte, error)
}
