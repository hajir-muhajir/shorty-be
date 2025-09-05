package service

import "testing"

func TestHasher_HashAndVerify(t *testing.T) {
	h := NewHasher()
	pw := "secret123"
	hash, err := h.Hash(pw)
	if err != nil{
		t.Fatalf("hash error: %v", err)
	}

	if hash == pw{
		t.Fatalf("hash should not equal password")
	}

	if !h.Verify(hash, pw){
		t.Fatalf("expected password to verify correctly")
	}

	if h.Verify(hash, "wrong password"){
		t.Fatalf("expected password verification to fail for wrong password")
	}
}