package service

import (
	"crypto/rand"
	"math/big"
	"regexp"
)

var base62 = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var aliasRE = regexp.MustCompile(`^[a-z0-9-]{3,30}$`)

type AliasGenerator struct{}

func NewAliasGenerator() *AliasGenerator {
	return &AliasGenerator{}
}

func (g *AliasGenerator) Generate(n int) string {
	if n < 6 {
		n = 6
	}
	out := make([]rune, n)
	for i := range out {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(len(base62))))
		out[i] = base62[j.Int64()]
	}
	return string(out)
}

func (g *AliasGenerator) ValidCustom(s string) bool {
	return aliasRE.MatchString(s)
}
