package porta_cipher

// ------------------------------------------------ ---------------------------------------------------------------------

// Table 加密时使用的表
type Table [][]rune

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
