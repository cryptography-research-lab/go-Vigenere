package vigenere

import "errors"

var ErrInputCharacter = errors.New("")

var ErrTableCharacterMustLetters = errors.New("")

var ErrTableRowCharacterNotUniq = errors.New("")

var ErrTableRowCount = errors.New("")
