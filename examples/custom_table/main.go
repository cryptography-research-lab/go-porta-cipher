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
