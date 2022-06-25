package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

func createHash(key string) []byte {
	// Generating 32 byte key
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

func Encrypt(data []byte, passphrase string) (string, error) {

	cipherBlock, err := aes.NewCipher(createHash(passphrase))
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aead.Seal(nonce, nonce, data, nil)

	return base64.URLEncoding.EncodeToString(ciphertext), nil

}

func Decrypt(data string, passphrase string) (string, error) {
	encryptData, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	cipherBlock, err := aes.NewCipher(createHash(passphrase))
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonceSize := aead.NonceSize()
	if len(encryptData) < nonceSize {
		return "", err
	}

	nonce, cipherText := encryptData[:nonceSize], encryptData[nonceSize:]
	plainData, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainData), nil
}

// func encryptFile(filename string, data []byte, passphrase string) error {
// 	f, _ := os.Create(filename)
// 	defer f.Close()
// 	cipher, err := Encrypt(data, passphrase)
// 	if err != nil {
// 		return err
// 	}
// 	f.Write(cipher)
// 	return nil
// }
//
// func decryptFile(filename string, passphrase string) ([]byte, error) {
// 	data, _ := ioutil.ReadFile(filename)
// 	plaintext, err := Decrypt(data, passphrase)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	return plaintext, nil
// }
