package crypto

import (
	"crypto/rand"
	"io"

	"github.com/pkg/errors"
	"golang.org/x/crypto/nacl/secretbox"
)

// Encrypter is an interface fulfilled by a provider to look up keys and encrypt the provided contents
type Encrypter interface {
	EncryptionKeyProvider
	Encrypt(contents []byte, dataEncryptionKey []byte) ([]byte, error)
}

// Encrypt a message using NaCl SecretBox
func Encrypt(contents []byte, key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, KeyError{}
	}
	var secretKey [32]byte
	copy(secretKey[:], key)

	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return nil, errors.Wrap(err, "could not construct random nonce")
	}

	// store encypted message as [24]byte nonce followed by encrypted file
	encrypted := secretbox.Seal(nonce[:], contents, &nonce, &secretKey)
	return encrypted, nil
}
