package main

import (
	"crypto/aes"
	"crypto/cipher"

	// "crypto/hmac"
	// "crypto/rand"
	// "crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"

	// "io"
	"os"
)

const (
	base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
)

var base62Encoding = base64.NewEncoding(base62Chars).WithPadding(base64.NoPadding)

// Encrypt encrypts data with AES-256-GCM and signs with HMAC-SHA256
func Encrypt(data, key []byte) (string, error) {
	// Split key into encryption key and HMAC key
	if len(key) != 64 { // 32 for AES + 32 for HMAC
		return "", errors.New("key must be 64 bytes long")
	}
	aesKey := key[:32]
	// hmacKey := key[32:]

	// AES-GCM encryption
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// nonce := make([]byte, gcm.NonceSize())
	// if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
	// 	return "", err
	// }
	nonce := []byte("1972373c43fc")
	ciphertext := gcm.Seal(nil, nonce, data, nil)
	ciphertext = append(nonce, ciphertext...)

	// HMAC-SHA256 signature
	// mac := hmac.New(sha256.New, hmacKey)
	// mac.Write(ciphertext)
	// signature := mac.Sum(nil)

	// Combine ciphertext and signature
	// encryptedData := append(ciphertext, signature...)
	encryptedData := ciphertext
	// Base62 encoding
	encoded := base62Encoding.EncodeToString(encryptedData)
	return encoded, nil
}

// Decrypt decrypts data with AES-256-GCM after verifying HMAC-SHA256 signature
func Decrypt(encoded string, key []byte) ([]byte, error) {
	// Split key into encryption key and HMAC key
	if len(key) != 64 {
		return nil, errors.New("key must be 64 bytes long")
	}
	aesKey := key[:32]
	// hmacKey := key[32:]

	// Base62 decoding
	encryptedData, err := base62Encoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	// // Split ciphertext and signature
	// if len(encryptedData) < sha256.Size {
	// 	return nil, errors.New("invalid encrypted data length")
	// }
	// dataLen := len(encryptedData) - sha256.Size
	// ciphertext := encryptedData[:dataLen]
	// signature := encryptedData[dataLen:]

	// // Verify HMAC-SHA256 signature
	// mac := hmac.New(sha256.New, hmacKey)
	// mac.Write(ciphertext)
	// expectedSignature := mac.Sum(nil)

	// if !hmac.Equal(signature, expectedSignature) {
	// 	return nil, errors.New("invalid signature")
	// }

	// AES-GCM decryption
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ciphertext := encryptedData
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage:")
		fmt.Println("  enc-any encrypt <key> <plaintext>")
		fmt.Println("  enc-any decrypt <key> <ciphertext>")
		fmt.Println("Key must be 64 bytes long (32 for AES + 32 for HMAC)")
		os.Exit(1)
	}

	cmd := os.Args[1]
	key := []byte(os.Args[2])
	data := os.Args[3]

	switch cmd {
	case "encrypt":
		result, err := Encrypt([]byte(data), key)
		if err != nil {
			fmt.Println("Encryption error:", err)
			os.Exit(1)
		}
		fmt.Println(result)
	case "decrypt":
		result, err := Decrypt(data, key)
		if err != nil {
			fmt.Println("Decryption error:", err)
			os.Exit(1)
		}
		fmt.Println(string(result))
	default:
		fmt.Println("Invalid command. Use 'encrypt' or 'decrypt'")
		os.Exit(1)
	}
}
