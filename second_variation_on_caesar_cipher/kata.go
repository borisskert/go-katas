package kata

/*
 * https://www.codewars.com/kata/55084d3898b323f0aa000546/train/go
 */

import (
	"math"
	"strings"
	"unicode"
)

func Encode(text string, shift int) []string {
	prefix := cleartext(text).prefix(shift)
	ciphertext := prefix.encrypt(text)

	return ciphertext.chunks(5)
}

func Decode(arr []string) string {
	ciphertext := chunks(arr).ciphertext()
	return string(ciphertext.decrypt())
}

type Prefix struct {
	prefix string
	shift  int
}

func (ctx Prefix) encrypt(cleartext string) ciphertext {
	encrypted := encrypt(cleartext, ctx.shift)
	return ciphertext(ctx.prefix + encrypted)
}

type cleartext string

func (ctx cleartext) prefix(shift int) Prefix {
	first := unicode.ToLower([]rune(ctx)[0])
	second := encryptRune(first, shift)

	return Prefix{string([]rune{first, second}), shift}
}

type ciphertext string

func (ctx ciphertext) chunkSize(amount int) int {
	textLength := len(ctx)
	return int(math.Ceil(float64(textLength) / float64(amount)))
}

func (ctx ciphertext) chunks(amount int) chunks {
	chunkSize := ctx.chunkSize(amount)
	return makeChunks(string(ctx), chunkSize)
}

const alphabetSize = 26

func (ctx ciphertext) shift() int {
	first := []rune(ctx)[0]
	second := []rune(ctx)[1]

	return (int(second) - int(first) + alphabetSize) % alphabetSize
}

func (ctx ciphertext) decrypt() cleartext {
	shift := ctx.shift()
	payload := string(ctx)[2:]
	decrypted := encrypt(payload, shift*-1)

	return cleartext(decrypted)
}

type chunks []string

func (ctx chunks) ciphertext() ciphertext {
	return ciphertext(strings.Join(ctx, ""))
}

func encrypt(cleartext string, shift int) string {
	var ciphertext = []rune{}

	for _, element := range cleartext {
		encrypted := encryptRune(element, shift)
		ciphertext = append(ciphertext, encrypted)
	}

	return string(ciphertext)
}

func encryptRune(cleartext rune, shift int) rune {
	var offset int8

	switch {
	case unicode.IsLower(cleartext):
		offset = int8('a')
	case unicode.IsUpper(cleartext):
		offset = int8('A')
	default:
		return cleartext
	}

	code := (int8(cleartext)-offset+int8(shift%alphabetSize)+alphabetSize)%alphabetSize + offset

	return rune(code)
}

func makeChunks(ciphertext string, chunkSize int) []string {
	if len(ciphertext) < chunkSize {
		return []string{ciphertext}
	}

	chunk := ciphertext[0:chunkSize]
	remaining := ciphertext[chunkSize:]

	if len(remaining) > 0 {
		furtherChunks := makeChunks(remaining, chunkSize)

		return append([]string{chunk}, furtherChunks...)
	}

	return []string{chunk}
}
