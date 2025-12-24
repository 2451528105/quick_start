package nrandom

import (
	"math/rand"
)

func GetRandomString(length int) string {
	if length <= 0 {
		return ""
	}
	// 字母集合（大写和小写）
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	// 字母和数字集合
	lettersAndDigits := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	result := make([]byte, length)

	// 第一个字符必须是字母
	result[0] = letters[rand.Intn(len(letters))]

	// 后续字符可以是字母或数字
	for i := 1; i < length; i++ {
		result[i] = lettersAndDigits[rand.Intn(len(lettersAndDigits))]
	}

	return string(result)
}
