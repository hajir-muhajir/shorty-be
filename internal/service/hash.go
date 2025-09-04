package service

import "golang.org/x/crypto/bcrypt"

type Hasher struct{}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (h *Hasher) Hash(pw string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(b), err
}

func (h *Hasher) Verify(hash, pw string)bool{
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw)) == nil
}