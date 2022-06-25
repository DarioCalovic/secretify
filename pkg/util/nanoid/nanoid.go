package nanoid

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const (
	alphabet = "346789ABCDEFGHJKLMNPQRTUVWXYabcdefghijkmnpqrtwxyz"
)

func GenerateIdentifier(size int) (string, error) {
	return gonanoid.Generate(alphabet, size)
}
