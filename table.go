package porta_cipher

// ------------------------------------------------ ---------------------------------------------------------------------

type Table [][]rune

// Query 根据行列的字母查询其对应的字母
func (x Table) Query(rowCharacter rune, columnCharacter rune) (rune, error) {
	rowCharacter = toUppercaseIfNeed(rowCharacter)
	columnCharacter = toUppercaseIfNeed(columnCharacter)
	if rowCharacter < 'A' || rowCharacter > 'Z' {
		return ' ', ErrInputError
	}
	if columnCharacter < 'A' || columnCharacter > 'Z' {
		return ' ', ErrInputError
	}
	rowIndex := (rowCharacter - 'A') / 2
	columnINdex := columnCharacter - 'A'
	return x[rowIndex][columnINdex], nil
}

func toUppercaseIfNeed(character rune) rune {
	if character >= 'a' && character <= 'z' {
		character -= 32
	}
	return character
}

// ------------------------------------------------ ---------------------------------------------------------------------

// DefaultTable 默认的表
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
