package util

const (
	symbols       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbolsLength = uint(len(symbols))
)

func EncodeBase62(number uint) string {
	if number == 0 {
		return string(symbols[0])
	}

	chars := make([]byte, 0, 7)

	for number > 0 {
		remainder := number % symbolsLength
		number = number / symbolsLength
		chars = append([]byte{symbols[remainder]}, chars...)
	}

	return string(chars)
}
