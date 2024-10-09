package core

import (
	"math/rand"
	"time"
)

type RandomTokenGenerator interface {
	Generate(tokenNumber int) (token string, err error)
}

type randomTokenGenerator struct {
}

var randomTokenGeneratorInstance RandomTokenGenerator

func GetRandomTokenGenerator() RandomTokenGenerator {
	if randomTokenGeneratorInstance == nil {
		randomTokenGeneratorInstance = newRandomTokenGenerator()
	}

	return randomTokenGeneratorInstance
}

func newRandomTokenGenerator() RandomTokenGenerator {
	return &randomTokenGenerator{}
}

func (rtg *randomTokenGenerator) Generate(tokenNumber int) (string, error) {
	rand.Seed(time.Now().Unix())
	length := tokenNumber

	ran_str := make([]byte, length)

	// Generating Random string
	for i := 0; i < length; i++ {
		ran_str[i] = byte(65 + rand.Intn(25))
	}

	return string(ran_str), nil
}
