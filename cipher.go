package shared

import (
	"encoding/base64"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
)

// Cipher is an interface for something that can encrypt and decrypt strings.
type Cipher interface {
	Decrypt(s string) (string, error)
	Encrypt(s string) (string, error)
}

// KMSCipher is a Cipher that performs symmetric encryption with an AWS KMS key.
// TODO(elffjs): Give this a proper constructor.
// TODO(elffjs): Figure out key rotation in a couple of weeks.
type KMSCipher struct {
	Client *kms.KMS
	KeyID  string
}

func (e *KMSCipher) Encrypt(s string) (string, error) {
	out, err := e.Client.Encrypt(
		&kms.EncryptInput{
			KeyId:               aws.String(e.KeyID),
			Plaintext:           []byte(s),
			EncryptionAlgorithm: aws.String(kms.EncryptionAlgorithmSpecSymmetricDefault),
		},
	)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(out.CiphertextBlob), nil
}

func (e *KMSCipher) Decrypt(s string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	out, err := e.Client.Decrypt(&kms.DecryptInput{
		CiphertextBlob:      b,
		KeyId:               aws.String(e.KeyID),
		EncryptionAlgorithm: aws.String(kms.EncryptionAlgorithmSpecSymmetricDefault),
	})
	if err != nil {
		return "", err
	}
	return string(out.Plaintext), nil
}

// ROT13Cipher is a Cipher that implements ROT13 encryption. Do not use this outside of tests!
type ROT13Cipher struct{}

func (e *ROT13Cipher) Decrypt(s string) (string, error) {
	return rot13(s), nil
}

func (e *ROT13Cipher) Encrypt(s string) (string, error) {
	return rot13(s), nil
}

func rot13(s string) string {
	return strings.Map(rot13Rune, s)
}

func rot13Rune(r rune) rune {
	switch {
	case 'A' <= r && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	case 'a' <= r && r <= 'z':
		return 'a' + (r-'a'+13)%26
	}
	return r
}
