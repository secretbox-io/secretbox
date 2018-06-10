package aws

import "github.com/secretbox-io/secretbox/crypto"

// Encrypt returns the NaCl secretbox encrypted ciphertext for contents.  This method delegates
// to the generic crypto.Encrypt method.
func (p AWSProvider) Encrypt(contents []byte, dataEncryptionKey []byte) ([]byte, error) {
	return crypto.Encrypt(contents, dataEncryptionKey)
}
