package util

import "math/rand"

// ランダムな任意の文字数の文字列の生成
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// ランダムな任意の桁数の整数の生成
func RandomInt(d int) int {
	if d <= 0 {
		return 0
	}
	min := 1
	for i := 1; i < d; i++ {
		min *= 10
	}
	max := min * 10
	return rand.Intn(max-min) + min
}
