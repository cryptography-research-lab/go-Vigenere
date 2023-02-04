package vigenere

import (
	"fmt"
	"testing"
)

func TestNewRandomTable(t *testing.T) {
}

func TestDefaultTable(t *testing.T) {
	//fmt.Println(DefaultTable.String())
}

func TestTable_TransformToDecrypt(t *testing.T) {
	decrypt := DefaultTable.TransformToDecrypt()
	fmt.Println(decrypt.String())
}
