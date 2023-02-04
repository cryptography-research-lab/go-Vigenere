package vigenere

import (
	"github.com/golang-infrastructure/go-shuffle"
	"strings"
)

// Table 用于表示维吉尼亚密码使用的密码表
type Table [][]rune

// NewRandomTable 创建一张随机的密码表用来在后面加密
func NewRandomTable() Table {
	table := make([][]rune, 26)
	for i := 0; i < 26; i++ {
		row := make([]rune, 26)
		for j := 0; j < 26; j++ {
			row[j] = rune('A' + j)
		}
		shuffle.Shuffle(row)
		table[i] = row
	}
	return table
}

// TransformToDecrypt 把表转换为适合解密查询的样子
func (x Table) TransformToDecrypt() Table {
	newTable := make([][]rune, len(x))
	for rowIndex, row := range x {
		newRow := make([]rune, 26)
		for index, character := range row {
			fromCharacter := rune('A' + index)
			newRow[character-'A'] = fromCharacter
		}
		newTable[rowIndex] = newRow
	}
	return newTable
}

// Query 根据行列的字母查询其对应的字母
func (x Table) Query(rowCharacter rune, columnCharacter rune) (rune, error) {

	// 先对输入做校验
	rowCharacter = toUppercaseIfNeed(rowCharacter)
	columnCharacter = toUppercaseIfNeed(columnCharacter)
	if rowCharacter < 'A' || rowCharacter > 'Z' {
		return ' ', ErrInputCharacter
	}
	if columnCharacter < 'A' || columnCharacter > 'Z' {
		return ' ', ErrInputCharacter
	}

	// 然后根据行列的字符做路由，找到对应的要映射到的字符
	rowIndex := rowCharacter - 'A'
	columnIndex := columnCharacter - 'A'
	return x[rowIndex][columnIndex], nil
}

// 校验这个密码表是否合法
func (x Table) check() error {
	if len(x) != 26 {
		return ErrTableRowCount
	}
	for _, row := range x {
		// 每个字母要恰好出现一次
		characterCount := make([]int, 26)
		for _, character := range row {
			character = toUppercaseIfNeed(character)
			if character < 'A' || character > 'Z' {
				return ErrTableCharacterMustLetters
			}
			// 统计出现次数
			characterCount[character-'A']++
		}
		// 检查统计的出现次数
		for _, count := range characterCount {
			if count != 1 {
				return ErrTableRowCharacterNotUniq
			}
		}
	}
	return nil
}

// 如果是小写字母的话，转为大写字母，如果是其它字符，则保持原样
func toUppercaseIfNeed(character rune) rune {
	if character >= 'a' && character <= 'z' {
		character -= 32
	}
	return character
}

// 把加密使用的表格转为字符串返回，用于观察表格长啥样
// 返回数据样例：
//
//	 [
//		[ I, C, L, O, M ]
//		[ P, H, D, R, Z ]
//		[ U, V, F, Y, B ]
//		[ G, X, T, Q, E ]
//		[ S, N, K, W, A ]
//	]
func (x Table) String() string {
	sb := strings.Builder{}
	sb.WriteString("[\n")
	for _, line := range x {
		sb.WriteString("\t[ ")
		for index, column := range line {
			sb.WriteRune(column)
			if index+1 != len(line) {
				sb.WriteString(",")
			}
			sb.WriteString(" ")
		}
		sb.WriteString("]\n")
	}
	sb.WriteString("]")
	return sb.String()
}

// ------------------------------------------------ ---------------------------------------------------------------------

// DefaultTable 维吉尼亚默认使用的密码表，这个密码表是约定俗成的
var DefaultTable Table

// 初始化默认密码表
func init() {
	DefaultTable = make([][]rune, 26)
	for i := 0; i < 26; i++ {
		row := make([]rune, 26)
		for j := 0; j < 26; j++ {
			row[j] = rune(((j + i) % 26) + 'A')
		}
		DefaultTable[i] = row
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------
