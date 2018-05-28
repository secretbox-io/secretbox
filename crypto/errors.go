package crypto

import "fmt"

// KeyError is returned when a key is malformed or missing
type KeyError struct{}

func (k KeyError) Error() string {
	return "malformed key or no key provided"
}

// KeyLookupError is returned when the key doesnt exist or an external key management service fails
type KeyLookupError struct {
	Msg string
}

func (k KeyLookupError) Error() string {
	switch len(k.Msg) {
	case 0:
		return fmt.Sprintf("key lookup error: %s", k.Msg)
	default:
		return "key lookup error"
	}
}

// DecryptionError is returned when an error is encountered with opening a SecretBox encoded message
type DecryptionError struct{}

func (d DecryptionError) Error() string {
	return "decryption failed"
}
