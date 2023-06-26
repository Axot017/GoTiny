package util

import "math/rand"

const (
	webSafeSymbols       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
	webSafeSymbolsLength = len(symbols)
)

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = webSafeSymbols[rand.Intn(webSafeSymbolsLength)]
	}
	return string(b)
}
