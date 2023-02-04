package vigenere

import (
	cycle_string "github.com/cryptography-research-lab/go-cycle-string"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
)

// Encrypt 对明文进行维吉尼亚密码加密
func Encrypt(plaintext, key string, table ...Table) (string, error) {

	// 设置默认参数
	table = variable_parameter.SetDefaultParam(table, DefaultTable)

	keyCycleString := cycle_string.NewCycleString(key)
	plaintextRuneSlice := []rune(plaintext)
	result := make([]rune, len(plaintextRuneSlice))
	for index, character := range plaintextRuneSlice {
		rowCharacter := keyCycleString.RuneAt(index)
		r, err := table[0].Query(rowCharacter, character)
		if err != nil {
			return "", err
		}
		result[index] = r
	}
	return string(result), nil
}

// Decrypt
// 对维吉尼亚加密的密文进行解密
func Decrypt(ciphertext, key string, table ...Table) (string, error) {
	table = variable_parameter.SetDefaultParam(table, DefaultTable)
	return Encrypt(ciphertext, key, table[0].TransformToDecrypt())
}
