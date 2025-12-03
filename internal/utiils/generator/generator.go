package generator

import gonanoid "github.com/matoous/go-nanoid/v2"

func GenCode(len int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return gonanoid.Generate(charset, len)
}

func GenCodeV2(charset string, len int) (string, error) {
	return gonanoid.Generate(charset, len)
}
