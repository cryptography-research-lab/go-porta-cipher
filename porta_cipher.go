package porta_cipher

import (
	cycle_string "github.com/cryptography-research-lab/go-cycle-string"
	"unicode/utf8"
)

// Encrypt 对明文使用Porta加密，输入需要全部是英文字母，大小写不区分
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

// Decrypt 对使用了Porta加密的密文进行解密
func Decrypt(ciphertext, key string) (string, error) {
	return Encrypt(ciphertext, key)
}
