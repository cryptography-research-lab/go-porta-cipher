package porta_cipher

import (
	cycle_string "github.com/cryptography-research-lab/go-cycle-string"
	"unicode/utf8"
)

// Encrypt 加密
func Encrypt(plaintext, key string) (string, error) {
	result := make([]rune, utf8.RuneCountInString(plaintext))
	keyString := cycle_string.NewCycleString(key)
	for index, character := range plaintext {
		rowCharacter := keyString.RuneAt(index)
		query, err := DefaultTable.Query(rowCharacter, character)
		if err != nil {
			return "", err
		}
		result[index] = query
	}
	return string(result), nil
}

// Decrypt 解密
func Decrypt(ciphertext, key string) (string, error) {
	return Encrypt(ciphertext, key)
}
