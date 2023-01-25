package porta_cipher

import (
	cycle_string "github.com/cryptography-research-lab/go-cycle-string"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
	"unicode/utf8"
)

// Encrypt 对明文使用Porta加密，输入需要全部是英文字母，大小写不区分
func Encrypt(plaintext, key string, table ...Table) (string, error) {

	// 设置默认值，如果没有手动制定加密表的话，则使用默认的
	variable_parameter.SetDefaultParam(table, DefaultTable)

	// 参数检查，即使是DefaultTable也要检查，因为它可能会被人为修改
	if err := table[0].check(); err != nil {
		return "", err
	}

	result := make([]rune, utf8.RuneCountInString(plaintext))
	keyString := cycle_string.NewCycleString(key)
	for index, character := range plaintext {
		rowCharacter := keyString.RuneAt(index)
		query, err := table[0].Query(rowCharacter, character)
		if err != nil {
			return "", err
		}
		result[index] = query
	}
	return string(result), nil
}

// Decrypt 对使用了Porta加密的密文进行解密
func Decrypt(ciphertext, key string, table ...Table) (string, error) {
	return Encrypt(ciphertext, key, table...)
}
