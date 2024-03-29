#  Porta加密(Porta Cipher)

# 一、安装

```bash
go get -u github.com/cryptography-research-lab/go-porta-cipher
```

# 二、使用示例

## 2.1 加密 & 解密示例代码
```go
package main

import (
	"fmt"
	porta_cipher "github.com/cryptography-research-lab/go-porta-cipher"
)

func main() {

	key := "THISISVERYSECURITYKEY"
	encrypt, err := porta_cipher.Encrypt("HELLOWORLD", key)
	if err != nil {·		fmt.Println("加密时发生了错误： " + err.Error())
		return
	}
	fmt.Println("加密结果： " + encrypt) // Output: 加密结果： QUPUKAECTP

	decrypt, err := porta_cipher.Decrypt(encrypt, key)
	if err != nil {
		fmt.Println("解密时发生了错误： " + err.Error())
		return
	}
	fmt.Println("解密结果： " + decrypt) // Output: 解密结果： HELLOWORLD

}
```
或者使用自定义的密码表，使用自定义密码表的时候需要注意保存密码表否则无法解密了：
```go
package main

import (
	"fmt"
	porta_cipher "github.com/cryptography-research-lab/go-porta-cipher"
)

func main() {

	// 使用自定义的密码表，可以用给定的默认方法生成一个随机的，当然密码表要保存好否则没办法解密了
	table := porta_cipher.NewRandomTable()
	fmt.Println("使用的密码表： ")
	fmt.Println(table.String())
	// Output:
	// 使用的密码表：
	//[
	//        [ T, W, R, H, X, F, O, D, S, U, Y, L, Z, Q, G, P, N, C, I, A, J, V, B, E, K, M ]
	//        [ I, Z, L, W, X, U, V, R, A, M, P, C, J, Q, S, K, N, H, O, T, F, G, D, E, Y, B ]
	//        [ M, Z, J, V, T, O, W, X, N, C, U, P, A, I, F, L, Q, S, R, E, K, D, G, H, Y, B ]
	//        [ B, A, S, M, H, J, O, E, Z, F, U, P, D, W, G, L, X, Y, C, T, K, V, N, Q, R, I ]
	//        [ X, N, W, S, Z, M, V, I, H, Q, P, T, F, B, U, K, J, Y, D, L, O, G, C, A, R, E ]
	//        [ D, F, P, A, I, B, L, V, E, X, O, G, Q, R, K, C, M, N, Y, Z, W, H, U, J, S, T ]
	//        [ D, E, I, A, B, Z, W, Y, C, N, R, Q, M, J, X, U, L, K, S, V, P, T, G, O, H, F ]
	//        [ W, F, T, N, J, B, V, L, U, E, R, H, X, D, O, Q, P, K, S, C, I, G, A, M, Z, Y ]
	//        [ N, B, S, M, O, W, K, H, Z, Y, G, V, D, A, E, P, Q, X, C, T, U, L, F, R, J, I ]
	//        [ R, G, X, I, V, T, B, K, D, W, H, N, M, L, Z, P, S, A, Q, F, U, E, J, C, Y, O ]
	//        [ C, Q, A, K, O, T, R, Y, M, P, D, V, I, X, E, J, B, G, Z, F, U, L, W, N, H, S ]
	//        [ D, I, Y, A, E, Q, J, N, B, G, X, W, Z, H, T, S, F, V, P, O, U, R, L, K, C, M ]
	//        [ N, U, V, P, O, Q, G, M, Z, J, R, Y, H, A, E, D, F, K, T, S, B, C, X, W, L, I ]
	//]

	key := "THISISVERYSECURITYKEY"
	encrypt, err := porta_cipher.Encrypt("HELLOWORLD", key, table)
	if err != nil {
		fmt.Println("加密时发生了错误： " + err.Error())
		return
	}
	// 不同的密码表得到的加密结果不同，这里只是一个例子
	fmt.Println("加密结果： " + encrypt) // Output: 加密结果： KHTNUJESVP

	decrypt, err := porta_cipher.Decrypt(encrypt, key, table)
	if err != nil {
		fmt.Println("解密时发生了错误： " + err.Error())
		return
	}
	fmt.Println("解密结果： " + decrypt) // Output: 解密结果： HELLOWORLD

}
```
## 2.2 破解代码示例

TODO 

# 三、Porta加密详解

## 3.1 加密

Porta类似于[维吉尼亚密码](https://github.com/cryptography-research-lab/go-Vigenere)，不同的是它的密码表只有13行，它有一个固定的密码表：

```text
KEYS| A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
----|----------------------------------------------------
A,B | N O P Q R S T U V W X Y Z A B C D E F G H I J K L M
C,D | O P Q R S T U V W X Y Z N M A B C D E F G H I J K L
E,F | P Q R S T U V W X Y Z N O L M A B C D E F G H I J K
G,H | Q R S T U V W X Y Z N O P K L M A B C D E F G H I J
I,J | R S T U V W X Y Z N O P Q J K L M A B C D E F G H I
K,L | S T U V W X Y Z N O P Q R I J K L M A B C D E F G H
M,N | T U V W X Y Z N O P Q R S H I J K L M A B C D E F G
O,P | U V W X Y Z N O P Q R S T G H I J K L M A B C D E F
Q,R | V W X Y Z N O P Q R S T U F G H I J K L M A B C D E
S,T | W X Y Z N O P Q R S T U V E F G H I J K L M A B C D
U,V | X Y Z N O P Q R S T U V W D E F G H I J K L M A B C
W,X | Y Z N O P Q R S T U V W X C D E F G H I J K L M A B
Y,Z | Z N O P Q R S T U V W X Y B C D E F G H I J K L M A
```

要求输入的要加密内容和秘钥都是英文字母，`x`坐标为明文字母对应的列，每一个明文字母对应一列，总共有26列，`y`坐标为秘钥字母对应的行，其中会出现两个字母对应一行的情况，所以总共有13行，因为处理的时候明文字母和秘钥字母是一一对应处理的，所以加密的时候需要先将秘钥重复自身直到字符长度和明文对齐，然后明文和秘钥相同下标对应的字母所对应的行和列的交点就是要映射到的字母，直接说有点绕不太好想明白，下面是一个实际的例子来说明加密的详细过程。

比如明文为`ABCD`，秘钥为`QWE`，则加密过程如下，最开始明文是这样子的：

![image-20230125231401207](README.assets/image-20230125231401207.png)

秘钥是这样子的：

![image-20230125231412282](README.assets/image-20230125231412282.png)

明文和秘钥的长度是不相等的：

![image-20230125231457146](README.assets/image-20230125231457146.png)

首先将密文`QWE`自身重复拼接直到与明文字符长度一样，这一步完成之后秘钥变为`QWEQ`，此时明文和秘钥的长度都是4：

![image-20230125231445134](README.assets/image-20230125231445134.png)

加密结果的长度和明文是相同的，现在的情况是这样，接下来要开始加密了： 

![image-20230125231507581](README.assets/image-20230125231507581.png)

然后看下标0，明文下标0的字母是A，秘钥的下标0的字母是Q，看加密表中A对应的列和Q对应的行的交点处的字母是V，所以加密结果的第一个字符就是V：

![image-20230125231515461](README.assets/image-20230125231515461.png)

然后看下标1，明文下标1的字母是B，秘钥下标1的字母W，交点处字母为Z：

![image-20230125231524097](README.assets/image-20230125231524097.png)

然后看下标2，明文下标2字母为C，秘钥下标2处为E，交点处字母为R：

![image-20230125231532451](README.assets/image-20230125231532451.png)

然后看下标3，明文下标3字母为D，秘钥下标3为Q，交点处字母为Y：

![image-20230125231539713](README.assets/image-20230125231539713.png)

此时得到了加密结果`VZRY`：

![image-20230125231230269](README.assets/image-20230125231230269.png)

## 3.2 解密

Porta加密的解密过程和加密过程是一样的，我们以加密后的密文`VZRY`和秘钥`QWE`进行一次解密，因为前面已经说过加密的过程了，所以这里从另一个角度来看解密的过程。

首先密文V是通过`A --> Q  --> V`得到的，解密的话是`V --> Q --> A `，让我们观察表中这两个映射关系，似乎有某些关系： 

![image-20230125232156738](README.assets/image-20230125232156738.png)

然后是`B --> W -- > Z`，解密`Z --> W --> B`： 

![image-20230125232643047](README.assets/image-20230125232643047.png)

然后是`C --> E --> R `，解密`R --> E --> C`： 

![image-20230125232818875](README.assets/image-20230125232818875.png)

然后是`D --> Q --> Y `，解密`Y --> Q --> D`：

![image-20230125232926661](README.assets/image-20230125232926661.png)

通过观察上面的解密过程，我们得出下面几个规律：

1. 秘钥只是用来决定使用13行中的哪一行的，除此之外就没有其他作用了
2. 将明文称为x，将密文称为x'，则只需要每一行满足下标x的值为x'，而x'的值为x即可，因此可以随机生成13对来填充，如此没必要非使用给定的密码表，只要满足定义自己定制也可以

经过验证后确定我的想法是可行的，我把它称之为对Porta的加密表的扩展，API已经增加了相应的支持。

# 四、Porta破解详解

TODO 

# 五、TODO

- 验证加密表是否可以是自定义的，会有什么影响 



