package porta_cipher

import (
	"fmt"
	"testing"
)

func TestNewRandomTable(t *testing.T) {
	table := NewRandomTable()
	fmt.Println(table)
}
