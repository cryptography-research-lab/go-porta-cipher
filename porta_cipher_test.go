package porta_cipher

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDecrypt(t *testing.T) {
	plaintext := "helloworld"
	key := "cryptographyisfuny"
	encrypt, err := Encrypt(plaintext, key)
	assert.Nil(t, err)
	assert.Equal(t, "VZXSFCLJYX", encrypt)

	decrypt, err := Decrypt(encrypt, key)
	assert.Nil(t, err)
	assert.Equal(t, strings.ToUpper(plaintext), decrypt)
}
