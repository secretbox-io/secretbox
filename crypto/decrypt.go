package crypto

import "golang.org/x/crypto/nacl/secretbox"

// Decrypter is fulfilled by a provider to look up data encryption keys and decrypt contents
type Decrypter interface {
	DecryptionKeyProvider
	Decrypt(encrypted []byte, dataEncryptionKey []byte) ([]byte, error)
}

// Decrypt an encrypted NaCl Secretbox encoded message
func Decrypt(encrypted []byte, key []byte) ([]byte, error) {
	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])

	var secretKey [32]byte
	copy(secretKey[:], key)

	decrypted, ok := secretbox.Open(nil, encrypted[24:], &decryptNonce, &secretKey)
	if !ok {
		return nil, DecryptionError{}
	}
	return decrypted, nil
}
