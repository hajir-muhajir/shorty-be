package service

import (
	"testing"
	"unicode/utf8"
)

// karakter yang diizinkan (harus base62)
const allowed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func isAllowed(s string) bool {
	// cek tiap rune hanya ada di 'allowed'
	for _, r := range s {
		found := false
		for _, ar := range allowed {
			if r == ar {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestGenerate_DefaultMinLength(t *testing.T){
	g := NewAliasGenerator()
	alias := g.Generate(0)
	if count:= utf8.RuneCountInString(alias); count != 6{
		t.Fatalf("expected length 6, got %d %s", count, alias)
	}
	if !isAllowed(alias){
		t.Fatalf("alias contains invalid character: %q", alias)
	}
}

func TestGenerate_CustomLength(t *testing.T){
	g := NewAliasGenerator()
	alias := g.Generate(8)
	if count:= utf8.RuneCountInString(alias); count != 8{
		t.Fatalf("expected length 8, got %d %s", count, alias)
	}
	if !isAllowed(alias){
		t.Fatalf("alias contains invalid character: %q", alias)
	}
}

func TestValidCustom(t *testing.T){
	g := NewAliasGenerator()

	valid := []string{
		"abc",
		"abc-123",
		"a-1-2",
		"abc123abc123abc123abc123abc123",
	}

	invalid := []string{
		"Abc",
		"ab",
		"a_",
		"has space",
		"abc$",
		"abc123abc123abc123abc123abc123xxxxxx",
	}

	for _, s := range valid{
		if !g.ValidCustom(s){
			t.Fatalf("expected valid alias, got invalid: %q", s)
		}
	}

	for _, s := range invalid{
		if g.ValidCustom(s){
			t.Fatalf("expected invalid alias, go valid: %q", s)
		}
	}
}