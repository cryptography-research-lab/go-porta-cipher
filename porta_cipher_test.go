package porta_cipher

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDecrypt(t *testing.T) {

	table := NewRandomTable()

	plaintext := "helloworld"
	key := "cryptographyisfuny"
	encrypt, err := Encrypt(plaintext, key, table)
	assert.Nil(t, err)
	t.Log(encrypt)
	//assert.Equal(t, "VZXSFCLJYX", encrypt)

	decrypt, err := Decrypt(encrypt, key, table)
	assert.Nil(t, err)
	assert.Equal(t, strings.ToUpper(plaintext), decrypt)
	t.Log(decrypt)
}
