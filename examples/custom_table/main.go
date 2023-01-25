package main

import (
	"fmt"
	porta_cipher "github.com/cryptography-research-lab/go-porta-cipher"
)

func main() {

	// 使用自定义的密码表，可以用给定的默认方法生成一个随机的，当然密码表要保存好否则没办法解密了
	table := porta_cipher.NewRandomTable()

	key := "THISISVERYSECURITYKEY"
	encrypt, err := porta_cipher.Encrypt("HELLOWORLD", key, table)
	if err != nil {
		fmt.Println("加密时发生了错误： " + err.Error())
		return
	}
	// 不同的密码表得到的加密结果不同，这里只是一个例子
	fmt.Println("加密结果： " + encrypt) // Output: 加密结果： LVCHSOWBUQ

	decrypt, err := porta_cipher.Decrypt(encrypt, key, table)
	if err != nil {
		fmt.Println("解密时发生了错误： " + err.Error())
		return
	}
	fmt.Println("解密结果： " + decrypt) // Output: 解密结果： HELLOWORLD

}
