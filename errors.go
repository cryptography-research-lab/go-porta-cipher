package porta_cipher

import "errors"

var (
	// ErrInputCharacter 输入的字符中有不合法的字符，只允许英文字母输入
	ErrInputCharacter = errors.New("all characters must be English letters")
)
