package porta_cipher

import (
	"math/rand"
	"strings"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// Table 加密时使用的表
type Table [][]rune

// NewRandomTable 创建一张随机的加密表
func NewRandomTable() Table {
	table := make([][]rune, 13)
	for rowIndex := 0; rowIndex < 13; rowIndex++ {

		// step 01. 生成一张A-Z的集合
		letterSet := make(map[rune]struct{}, 0)
		for i := 0; i < 26; i++ {
			letterSet[rune('A'+i)] = struct{}{}
		}

		// step 02. 然后生成一行，就是不断配对，配对时保证不遗漏的同时有一些随机性
		row := make([]rune, 26)
		for i := 0; i < 26; i++ {
			letter := rune('A' + i)
			// 如果已经处理过了则不重复处理
			if _, exists := letterSet[letter]; !exists {
				continue
			}
			// set中随机获取一个
			//var matchTo rune
			//for key := range letterSet {
			//	matchTo = key
			//	break
			//}
			// 不借助语言特性了，慢点就慢点
			var matchTo rune
			index := rand.Intn(len(letterSet))
			for key := range letterSet {
				if index == 0 {
					matchTo = key
					break
				}
				index--
			}
			row[i] = matchTo
			row[matchTo-'A'] = letter
			delete(letterSet, matchTo)
			delete(letterSet, letter)
		}

		// step 03. 保存生成的行
		table[rowIndex] = row
	}
	return table
}

// 把加密使用的表格转为字符串返回，用于观察表格长啥样
// 返回数据样例：
//
//	 [
//		[ I, C, L, O, M ]
//		[ P, H, D, R, Z ]
//		[ U, V, F, Y, B ]
//		[ G, X, T, Q, E ]
//		[ S, N, K, W, A ]
//	]
func (x Table) String() string {
	sb := strings.Builder{}
	sb.WriteString("[\n")
	for _, line := range x {
		sb.WriteString("\t[ ")
		for index, column := range line {
			sb.WriteRune(column)
			if index+1 != len(line) {
				sb.WriteString(",")
			}
			sb.WriteString(" ")
		}
		sb.WriteString("]\n")
	}
	sb.WriteString("]")
	return sb.String()
}

// 校验这个密码表是否合法
func (x Table) check() error {
	for _, row := range x {
		// 每个字母要恰好出现一次，并且每一对的对应关系是OK的
		characterCount := make([]int, 26)
		for index, character := range row {
			character = toUppercaseIfNeed(character)
			if character < 'A' || character > 'Z' {
				return ErrTableCharacterMustLetters
			}
			// 统计出现次数
			characterCount[character-'A']++
			// 检查对应关系
			matchToIndex := character - 'A'
			if row[matchToIndex] != rune('A'+index) {
				return ErrTableMatchRelationNotOk
			}
		}
		// 检查统计的出现次数
		for _, count := range characterCount {
			if count != 1 {
				return ErrTableRowCharacterNotUniq
			}
		}
	}
	return nil
}

// Query 根据行列的字母查询其对应的字母
func (x Table) Query(rowCharacter rune, columnCharacter rune) (rune, error) {

	// 先对输入做校验
	rowCharacter = toUppercaseIfNeed(rowCharacter)
	columnCharacter = toUppercaseIfNeed(columnCharacter)
	if rowCharacter < 'A' || rowCharacter > 'Z' {
		return ' ', ErrInputCharacter
	}
	if columnCharacter < 'A' || columnCharacter > 'Z' {
		return ' ', ErrInputCharacter
	}

	// 然后根据行列的字符做路由，找到对应的要映射到的字符
	rowIndex := (rowCharacter - 'A') / 2
	columnIndex := columnCharacter - 'A'
	return x[rowIndex][columnIndex], nil
}

// 如果是小写字母的话，将其转为大写字母，否则将其原样返回
func toUppercaseIfNeed(character rune) rune {
	if character >= 'a' && character <= 'z' {
		character -= 32
	}
	return character
}

// ------------------------------------------------ ---------------------------------------------------------------------

// DefaultTable 默认的加密时使用的映射表，这张表是约定好的
// KEYS| A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
// ----|----------------------------------------------------
// A,B | N O P Q R S T U V W X Y Z A B C D E F G H I J K L M
// C,D | O P Q R S T U V W X Y Z N M A B C D E F G H I J K L
// E,F | P Q R S T U V W X Y Z N O L M A B C D E F G H I J K
// G,H | Q R S T U V W X Y Z N O P K L M A B C D E F G H I J
// I,J | R S T U V W X Y Z N O P Q J K L M A B C D E F G H I
// K,L | S T U V W X Y Z N O P Q R I J K L M A B C D E F G H
// M,N | T U V W X Y Z N O P Q R S H I J K L M A B C D E F G
// O,P | U V W X Y Z N O P Q R S T G H I J K L M A B C D E F
// Q,R | V W X Y Z N O P Q R S T U F G H I J K L M A B C D E
// S,T | W X Y Z N O P Q R S T U V E F G H I J K L M A B C D
// U,V | X Y Z N O P Q R S T U V W D E F G H I J K L M A B C
// W,X | Y Z N O P Q R S T U V W X C D E F G H I J K L M A B
// Y,Z | Z N O P Q R S T U V W X Y B C D E F G H I J K L M A
var DefaultTable Table = [][]rune{
	{'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M'},
	{'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'N', 'M', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L'},
	{'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'N', 'O', 'L', 'M', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K'},
	{'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'N', 'O', 'P', 'K', 'L', 'M', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'},
	{'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'N', 'O', 'P', 'Q', 'J', 'K', 'L', 'M', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'},
	{'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'N', 'O', 'P', 'Q', 'R', 'I', 'J', 'K', 'L', 'M', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'},
	{'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'N', 'O', 'P', 'Q', 'R', 'S', 'H', 'I', 'J', 'K', 'L', 'M', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
	{'U', 'V', 'W', 'X', 'Y', 'Z', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'A', 'B', 'C', 'D', 'E', 'F'},
	{'V', 'W', 'X', 'Y', 'Z', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'A', 'B', 'C', 'D', 'E'},
	{'W', 'X', 'Y', 'Z', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'A', 'B', 'C', 'D'},
	{'X', 'Y', 'Z', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'A', 'B', 'C'},
	{'Y', 'Z', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'A', 'B'},
	{'Z', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'A'},
}

// ------------------------------------------------ ---------------------------------------------------------------------
