package porta_cipher

import "errors"

// 加密表相关的错误
var (
	// ErrTableMatchRelationNotOk 加密表的对应关系不对
	ErrTableMatchRelationNotOk = errors.New("table match relation not ok")

	// ErrTableCharacterMustLetters 加密表之间的字符必须都是字母，可以是大写字母或者小写字母
	ErrTableCharacterMustLetters = errors.New("table character letters must letters")

	// ErrTableRowCharacterNotUniq ·加密表每行的每个字符仅能出现一次
	ErrTableRowCharacterNotUniq = errors.New("table row character not uniq ")
)

// 加密或解密的时候输入的key或者明文h或者密文不合法的错误
var (
	// ErrInputCharacter 输入的字符中有不合法的字符，只允许英文字母输入
	ErrInputCharacter = errors.New("all characters must be English letters")
)
